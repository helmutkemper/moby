// +build !exclude_graphdriver_zfs,linux !exclude_graphdriver_zfs,freebsd

package register // import "github.com/helmutkemper/moby/daemon/graphdriver/register"

import (
	// register the zfs driver
	_ "github.com/helmutkemper/moby/daemon/graphdriver/zfs"
)
