package main

import (
	"context"
	"fmt"
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
