package main

import (
	"context"
	"log"
	"strconv"
)

func handlerBrowse(s *state, cmd command) error {
	var limit int32 = 2
	if len(cmd.Args) == 1 {
		tempLimit, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			log.Printf("Not number")
			return err
		}
		limit = int32(tempLimit)
	}
	posts, err := s.db.GetPosts(context.Background(), limit)
	if err != nil {
		log.Printf("Failed to retrieve posts: %v", err)
		return err
	}
	log.Printf("Retrieved posts: %v", posts)
	for _, post := range posts {
		log.Printf("%+v\n", post)
	}
	return nil
}
