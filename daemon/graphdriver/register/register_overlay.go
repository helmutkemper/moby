// +build !exclude_graphdriver_overlay,linux

package register // import "github.com/helmutkemper/moby/daemon/graphdriver/register"

import (
	// register the overlay graphdriver
	_ "github.com/helmutkemper/moby/daemon/graphdriver/overlay"
)
