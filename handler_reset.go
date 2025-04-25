package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Things did not go as planned...: %w", err)
	}

	fmt.Println("Database has been reset!")
	return nil
}
