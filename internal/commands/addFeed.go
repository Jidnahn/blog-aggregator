package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/Jidnahn/blog-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerAddFeed(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("error: command expects url argument")
	}
	// get props form args
	name := user.Name
	url := cmd.Args[0]
	// create feed
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
	// create follow for feed
	follow, err := s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UpdatedAt: time.Now(),
		FeedID:    feed.ID,
		UserID:    feed.UserID,
	})
	if err != nil {
		return fmt.Errorf("error creating follow for created feed: %w", err)
	}

	fmt.Println(feed)
	fmt.Println(follow)

	return nil
}
