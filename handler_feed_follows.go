package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/samirhembrom/blogaggregator/internal/database"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("Not enough args")
	}

	url := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error fetching user %w", err)
	}
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error fetching feed %w", err)
	}

	feed_follows, err := s.db.CreateFeedFollow(
		context.Background(),
		database.CreateFeedFollowParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			UserID:    user.ID,
			FeedID:    feed.ID,
		},
	)
	if err != nil {
		return fmt.Errorf("error creating feed follow %w", err)
	}

	fmt.Println("Feed follow created:")
	printFeedFollow(feed_follows.UserName, feed_follows.FeedName)

	return nil
}

func handlerListFeedFollows(s *state, cmd command) error {
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("error fetching user %w", err)
	}

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error fetching feed %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Printf("Not following any feeds %d", len(feedFollows))
		return nil
	}

	fmt.Printf("Feed follows for user %s:\n", user.Name)
	for _, ff := range feedFollows {
		fmt.Printf("* %s\n", ff.FeedName)
	}

	return nil
}

func printFeedFollow(username, feedname string) {
	fmt.Printf("* User:          %s\n", username)
	fmt.Printf("* Feed:          %s\n", feedname)
}
