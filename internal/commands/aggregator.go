package commands

import (
	"context"
	"fmt"

	"github.com/Jidnahn/blog-aggregator/internal/rss"
)

func HanlderAgg(s *State, cmd Command) error {
	if len(cmd.Args) > 0 {
		return fmt.Errorf("error: command expects no arguments")
	}

	url := "https://www.wagslane.dev/index.xml"
	feed, err := rss.FetchFeed(context.Background(), url)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}
