package spotify

import (
	"context"

	zspotify "github.com/zmb3/spotify/v2"
)

func (c *Client) GetTrackByID(
	ctx context.Context,
	id string,
) (*zspotify.FullTrack, error) {

	return c.Client.GetTrack(
		ctx,
		zspotify.ID(id),
	)
}
