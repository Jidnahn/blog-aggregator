package commands

import (
	"context"
	"fmt"
	"time"

	"github.com/Jidnahn/blog-aggregator/internal/rss"
)

func HanlderAgg(s *State, cmd Command) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("error: command expects only one argument for time between requests")
	}
	// get time between requests
	time_between_reqs := cmd.Args[0]
	// define scape feeds function
	scrapeFeeds := func(s *State) error {
		// get feed to fetch
		feedToFetch, err := s.Db.GetNextFeedToFetch(context.Background())
		if err != nil {
			return fmt.Errorf("error retriving next feed to fetch: %w", err)
		}
		// mark feed as fetched
		fetched, err := s.Db.MarkFeedFetched(context.Background(), feedToFetch.ID)
		if err != nil {
			return fmt.Errorf("error fetching feed %s: %w", feedToFetch.ID, err)
		}
		// get the feed
		feed, err := rss.FetchFeed(context.Background(), fetched.Url)
		if err != nil {
			return err
		}
		// print the feed
		for _, item := range feed.Channel.Item {
			fmt.Println(item.Title)
			fmt.Println(item.Description)
			fmt.Println(item.Link)
			fmt.Println(item.PubDate)
			fmt.Println("-----")

		}
		return nil
	}
	// get interval for ticker
	interval, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("error parsing time between requests: %w", err)
	}
	// set ticker for scapeFeeds loop
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for ; ; <-ticker.C {
		err := scrapeFeeds(s)
		if err != nil {
			return err
		}
	}
}
