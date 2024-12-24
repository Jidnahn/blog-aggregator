package commands

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Jidnahn/blog-aggregator/internal/database"
)

func HandlerBrowse(s *State, cmd Command, user database.User) error {
	if len(cmd.Args) > 1 {
		return fmt.Errorf("error: command expects maximum one limit argument")
	}

	limit := 2
	if len(cmd.Args) == 1 {
		val, err := strconv.ParseInt(cmd.Args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("error parsing limit: %w", err)
		}
		if val <= 0 {
			return fmt.Errorf("error: limit must be a positive integer")
		}
		limit = int(val)
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error getting posts for user: %w", err)
	}

	if len(posts) == 0 {
		fmt.Println("No posts found")
		return nil
	}

	for _, post := range posts {
		fmt.Println(post.Title)
		fmt.Println(post.Description)
		fmt.Println(post.Url)
		fmt.Println(post.PublishedAt)
		fmt.Println("--------")
	}

	return nil
}
