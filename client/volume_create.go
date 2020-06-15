package client // import "github.com/helmutkemper/moby/client"

import (
	"context"
	"encoding/json"

	"github.com/helmutkemper/moby/api/types"
	volumetypes "github.com/helmutkemper/moby/api/types/volume"
)

// VolumeCreate creates a volume in the docker host.
func (cli *Client) VolumeCreate(ctx context.Context, options volumetypes.VolumeCreateBody) (types.Volume, error) {
	var volume types.Volume
	resp, err := cli.post(ctx, "/volumes/create", nil, options, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return volume, err
	}
	err = json.NewDecoder(resp.body).Decode(&volume)
	return volume, err
}
