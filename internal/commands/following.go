package commands

import (
	"context"
	"fmt"

	"github.com/Jidnahn/blog-aggregator/internal/database"
)

func HandlerFollowing(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("error: command expects no arguments")
	}
	// get follows that match the user ID
	follows, err := s.Db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("error getting followed feeds from current user: %w", err)
	}
	// print every follow in the console
	for _, feed := range follows {
		fmt.Println("Feed name:", feed.FeedName)
		fmt.Println("User name:", feed.UserName)
		fmt.Println("------")
	}

	return nil
}
