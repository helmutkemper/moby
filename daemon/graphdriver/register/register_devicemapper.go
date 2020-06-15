// +build !exclude_graphdriver_devicemapper,!static_build,linux

package register // import "github.com/helmutkemper/moby/daemon/graphdriver/register"

import (
	// register the devmapper graphdriver
	_ "github.com/helmutkemper/moby/daemon/graphdriver/devmapper"
)
