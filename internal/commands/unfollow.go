package commands

import (
	"context"
	"fmt"

	"github.com/Jidnahn/blog-aggregator/internal/database"
)

func HandlerUnfollow(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("error: command only takes one url argument")
	}

	url := cmd.Args[0]

	feed, err := s.Db.GetFeedByURL(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error getting feed from url: %w", err)
	}

	deleted, err := s.Db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		FeedID: feed.ID,
		UserID: user.ID,
	})
	if err != nil {
		return fmt.Errorf("error unfollowing feed: %w", err)
	}

	fmt.Println(deleted)

	return nil
}
