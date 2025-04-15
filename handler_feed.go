package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/samirhembrom/blogaggregator/internal/database"
)

func handlerListFeeds(s *state, _ command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	if len(feeds) == 0 {
		fmt.Printf("No records")
		return nil
	}

	fmt.Printf("Found %d feeds:\n", len(feeds))
	for _, feed := range feeds {
		user, err := s.db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return err
		}
		printFeed(feed, user)
	}

	return nil
}

func handlerAddFeed(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("Current user not found %w", err)
	}

	if len(cmd.Args) != 2 {
		return fmt.Errorf("Not enough args")
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

	fmt.Println("Feed created successfull:")
	printFeed(feed, user)
	return nil
}

func printFeed(feed database.Feed, user database.User) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
	fmt.Printf("* User:          %s\n", user.Name)
}
