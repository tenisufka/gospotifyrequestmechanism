package lyrics

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Client struct {
	httpClient *http.Client
}

func New() *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *Client) GetLyrics(
	ctx context.Context,
	artist string,
	title string,
) (string, error) {

	endpoint :=
		"https://lrclib.net/api/get" +
			"?artist_name=" + url.QueryEscape(artist) +
			"&track_name=" + url.QueryEscape(title)

	var lastErr error

	for attempt := 1; attempt <= 3; attempt++ {

		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodGet,
			endpoint,
			nil,
		)
		if err != nil {
			return "", err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			lastErr = err
			time.Sleep(time.Second)
			continue
		}

		var result LyricsResponse

		func() {
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				lastErr = fmt.Errorf(
					"lrclib returned %d",
					resp.StatusCode,
				)
				return
			}

			if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
				lastErr = err
				return
			}

			if result.PlainLyrics == "" {
				lastErr = fmt.Errorf("lyrics not found")
				return
			}

			lastErr = nil
		}()

		if lastErr == nil {
			return result.PlainLyrics, nil
		}

		time.Sleep(time.Second)
	}

	return "", lastErr
}
