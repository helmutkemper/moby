// +build !linux,!darwin,!freebsd,!windows

package daemon // import "github.com/helmutkemper/moby/daemon"

func (daemon *Daemon) setupDumpStackTrap(_ string) {
	return
}
