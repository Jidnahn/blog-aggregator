package commands

import (
	"context"
	"fmt"
)

func HandlerUsers(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("error: commands takes no arguments")
	}

	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users: %w", err)
	}
	currentUser := s.Config.Current_user_name
	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)", user.Name)
		} else {
			fmt.Println("*", user.Name)
		}
	}

	return nil
}
