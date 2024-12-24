package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/Jidnahn/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("error: command expects a single url argument")
	}
	// get prop from args
	url := cmd.Args[0]
	feed, err := s.Db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed by url: %w", err)
	}
	// create feed follow
	follow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("error following feed: %w", err)
	}

	fmt.Println(follow)

	return nil
}
