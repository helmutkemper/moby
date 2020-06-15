package daemon // import "github.com/helmutkemper/moby/daemon"

import (
	"github.com/helmutkemper/moby/container"
	"github.com/helmutkemper/moby/pkg/archive"
)

func (daemon *Daemon) tarCopyOptions(container *container.Container, noOverwriteDirNonDir bool) (*archive.TarOptions, error) {
	return daemon.defaultTarCopyOptions(noOverwriteDirNonDir), nil
}
