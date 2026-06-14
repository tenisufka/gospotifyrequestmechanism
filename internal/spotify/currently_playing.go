package spotify

import (
	"context"

	zspotify "github.com/zmb3/spotify/v2"
)

func (c *Client) CurrentlyPlaying(
	ctx context.Context,
) (*zspotify.FullTrack, error) {

	player, err := c.Client.PlayerCurrentlyPlaying(ctx)

	if err != nil {
		return nil, err
	}

	if player == nil || player.Item == nil {
		return nil, nil
	}

	return player.Item, nil
}
