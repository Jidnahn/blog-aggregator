package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/Jidnahn/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("error: command expects name and url arguments")
	}

	name := cmd.Args[0]
	url := cmd.Args[1]

	user, err := s.Db.GetUser(context.Background(), s.Config.Current_user_name)
	if err != nil {
		return fmt.Errorf("error getting current user: %w", err)
	}

	feed, err := s.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error creating feed: %w", err)
	}

	fmt.Println(feed)

	return nil
}
