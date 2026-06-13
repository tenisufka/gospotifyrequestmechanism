package spotify

import (
	"context"

	zspotify "github.com/zmb3/spotify/v2"
)

func (c *Client) AddToQueue(
	ctx context.Context,
	trackID zspotify.ID,
	deviceID string,
) error {

	return c.Client.QueueSong(
		ctx,
		trackID,
	)
}
