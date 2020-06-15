// +build !exclude_graphdriver_btrfs,linux

package register // import "github.com/helmutkemper/moby/daemon/graphdriver/register"

import (
	// register the btrfs graphdriver
	_ "github.com/helmutkemper/moby/daemon/graphdriver/btrfs"
)
