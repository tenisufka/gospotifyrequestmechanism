package spotify

import (
	"context"
	"fmt"

	zspotify "github.com/zmb3/spotify/v2"
)

func (c *Client) GetActiveDevice(
	ctx context.Context,
) (*zspotify.PlayerDevice, error) {

	devices, err := c.Client.PlayerDevices(ctx)

	if err != nil {
		return nil, err
	}

	for _, d := range devices {

		if d.Active {
			return &d, nil
		}
	}

	return nil, fmt.Errorf(
		"no active device found",
	)
}
