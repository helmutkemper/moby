// +build !linux,!freebsd,!windows

package daemon // import "github.com/helmutkemper/moby/daemon"

import (
	"github.com/helmutkemper/moby/daemon/config"
	"github.com/helmutkemper/moby/pkg/sysinfo"
)

const platformSupported = false

func setupResolvConf(config *config.Config) {
}

// RawSysInfo returns *sysinfo.SysInfo .
func (daemon *Daemon) RawSysInfo(quiet bool) *sysinfo.SysInfo {
	return sysinfo.New(quiet)
}
