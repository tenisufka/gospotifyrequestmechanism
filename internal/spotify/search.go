package spotify

import (
	"context"
	"errors"

	zspotify "github.com/zmb3/spotify/v2"
)

var ErrTrackNotFound = errors.New("track not found")

func (c *Client) SearchTrack(
	ctx context.Context,
	query string,
) (*zspotify.FullTrack, error) {

	result, err := c.Client.Search(
		ctx,
		query,
		zspotify.SearchTypeTrack,
	)

	if err != nil {
		return nil, err
	}

	if result.Tracks == nil ||
		len(result.Tracks.Tracks) == 0 {

		return nil, ErrTrackNotFound
	}

	id := result.Tracks.Tracks[0].ID

	return c.Client.GetTrack(
		ctx,
		id,
	)
}
