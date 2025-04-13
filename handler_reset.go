package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, _ command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error deleteing rows: %w", err)
	}

	fmt.Printf("Reset was successful")
	return nil
}
