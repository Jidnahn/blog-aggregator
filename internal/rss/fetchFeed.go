package rss

import (
	"context"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
)

func FetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	// create request
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("User-Agent", "gator")
	// get client
	client := &http.Client{}
	// make request
	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()
	// check OK status code
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}
	// decode response
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}
	var rssFeed RSSFeed
	if err := xml.Unmarshal(data, &rssFeed); err != nil {
		return nil, fmt.Errorf("error decoding data: %w", err)
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i, item := range rssFeed.Channel.Item {
		rssFeed.Channel.Item[i].Title = html.UnescapeString(item.Title)
		rssFeed.Channel.Item[i].Description = html.UnescapeString(item.Description)
	}

	return &rssFeed, nil
}
