// +build !exclude_graphdriver_fuseoverlayfs,linux

package register // import "github.com/helmutkemper/moby/daemon/graphdriver/register"

import (
	// register the fuse-overlayfs graphdriver
	_ "github.com/helmutkemper/moby/daemon/graphdriver/fuse-overlayfs"
)
