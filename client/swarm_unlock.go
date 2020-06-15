package client // import "github.com/helmutkemper/moby/client"

import (
	"context"

	"github.com/helmutkemper/moby/api/types/swarm"
)

// SwarmUnlock unlocks locked swarm.
func (cli *Client) SwarmUnlock(ctx context.Context, req swarm.UnlockRequest) error {
	serverResp, err := cli.post(ctx, "/swarm/unlock", nil, req, nil)
	ensureReaderClosed(serverResp)
	return err
}
