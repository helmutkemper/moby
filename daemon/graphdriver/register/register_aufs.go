// +build !exclude_graphdriver_aufs,linux

package register // import "github.com/helmutkemper/moby/daemon/graphdriver/register"

import (
	// register the aufs graphdriver
	_ "github.com/helmutkemper/moby/daemon/graphdriver/aufs"
)
