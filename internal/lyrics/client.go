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
		"https://api.lyrics.ovh/v1/" +
			url.PathEscape(artist) +
			"/" +
			url.PathEscape(title)

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
		return "", err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {

		return "", fmt.Errorf(
			"lyrics api returned %d",
			resp.StatusCode,
		)
	}

	var result LyricsResponse

	if err := json.NewDecoder(
		resp.Body,
	).Decode(&result); err != nil {

		return "", err
	}

	if result.Lyrics == "" {
		return "", fmt.Errorf(
			"lyrics not found",
		)
	}

	return result.Lyrics, nil
}
