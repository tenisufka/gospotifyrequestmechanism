package spotify

import (
	"context"
	"errors"

	zspotify "github.com/zmb3/spotify/v2"
)

var ErrNoDevice = errors.New(
	"no active spotify device",
)

func (c *Client) ActiveDevice(
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

	if len(devices) > 0 {
		return &devices[0], nil
	}

	return nil, ErrNoDevice
}
