package daemon // import "github.com/helmutkemper/moby/daemon"

import (
	"github.com/helmutkemper/moby/api/types/container"
	libcontainerdtypes "github.com/helmutkemper/moby/libcontainerd/types"
)

func toContainerdResources(resources container.Resources) *libcontainerdtypes.Resources {
	// We don't support update, so do nothing
	return nil
}
