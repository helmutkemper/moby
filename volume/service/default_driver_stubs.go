// +build !linux,!windows

package service // import "github.com/helmutkemper/moby/volume/service"

import (
	"github.com/helmutkemper/moby/pkg/idtools"
	"github.com/helmutkemper/moby/volume/drivers"
)

func setupDefaultDriver(_ *drivers.Store, _ string, _ idtools.Identity) error { return nil }
