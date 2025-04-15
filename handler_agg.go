package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/samirhembrom/blogaggregator/internal/database"
)

func handlerAgg(s *state, cmd command) error {
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error %w", err)
	}
	fmt.Printf("Feed: %+v\n", feed)
	return nil
}

func handlerFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("Not enough args")
	}

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Current user not found %w", err)
	}

	feedName := cmd.Args[0]
	feedUrl := cmd.Args[1]

	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedUrl,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error created feed: %w", err)
	}

	fmt.Printf("Name: %s\nUrl: %s\nUserId: %s", feed.Name, feed.Url, feed.UserID)
	return nil
}
