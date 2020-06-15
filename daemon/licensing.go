package daemon // import "github.com/helmutkemper/moby/daemon"

import (
	"github.com/helmutkemper/moby/api/types"
	"github.com/helmutkemper/moby/dockerversion"
)

func (daemon *Daemon) fillLicense(v *types.Info) {
	v.ProductLicense = dockerversion.DefaultProductLicense
}
