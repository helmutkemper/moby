package checkpoint // import "github.com/helmutkemper/moby/api/server/router/checkpoint"

import "github.com/helmutkemper/moby/api/types"

// Backend for Checkpoint
type Backend interface {
	CheckpointCreate(container string, config types.CheckpointCreateOptions) error
	CheckpointDelete(container string, config types.CheckpointDeleteOptions) error
	CheckpointList(container string, config types.CheckpointListOptions) ([]types.Checkpoint, error)
}
