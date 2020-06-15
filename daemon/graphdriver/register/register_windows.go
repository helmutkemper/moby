package register // import "github.com/helmutkemper/moby/daemon/graphdriver/register"

import (
	// register the windows graph drivers
	_ "github.com/helmutkemper/moby/daemon/graphdriver/lcow"
	_ "github.com/helmutkemper/moby/daemon/graphdriver/windows"
)
