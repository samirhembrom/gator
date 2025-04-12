package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/samirhembrom/blogaggregator/internal/database"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	existingUser, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		return fmt.Errorf("a user with the name '%s' already exists %w", name, err)
	}
	err = s.cfg.SetUser(existingUser.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("Username switched successful")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.Name)
	}

	name := cmd.Args[0]

	userData := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}
	user, err := s.db.CreateUser(context.Background(), userData)
	if err != nil {
		return fmt.Errorf("alread a user: %w", err)
	}

	err = s.cfg.SetUser(name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("Username created successful")
	fmt.Printf("Username switched successful %s\n", user.Name)
	return nil
}
