package commands

import (
	"context"
	"fmt"
	"strings"
)

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("error: expected at least one argument")
	}

	userName := strings.Join(cmd.Args, " ")
	user, err := s.Db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("error getting user: %w", err)
	}

	if err := s.Config.SetUser(userName); err != nil {
		return fmt.Errorf("error setting user in config: %w", err)
	}

	fmt.Printf("Logged in successfully with username %s\n", userName)
	fmt.Println(user)

	return nil
}
