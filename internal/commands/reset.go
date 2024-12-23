package commands

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("error: command takes no arguments")
	}

	if err := s.Db.DeleteAllUsers(context.Background()); err != nil {
		return fmt.Errorf("error deleting data from table users: %w", err)
	}

	return nil
}
