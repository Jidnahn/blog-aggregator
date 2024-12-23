package commands

import (
	"context"
	"fmt"
)

func HandlerFeeds(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("error: command expects no arguments")
	}

	feeds, err := s.Db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("error getting feeds from db: %w", err)
	}

	for _, feed := range feeds {
		user, err := s.Db.GetUserById(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("error getting user name from feed %d: %w", feed.ID, err)
		}
		fmt.Println("Feed", feed.ID)
		fmt.Println("Name:", feed.Name)
		fmt.Println("URL:", feed.Url)
		fmt.Println("Username:", user.Name)
		fmt.Println("-----")
	}

	return nil
}
