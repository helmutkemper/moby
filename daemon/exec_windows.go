package daemon // import "github.com/helmutkemper/moby/daemon"

import (
	"github.com/helmutkemper/moby/container"
	"github.com/helmutkemper/moby/daemon/exec"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

func (daemon *Daemon) execSetPlatformOpt(c *container.Container, ec *exec.Config, p *specs.Process) error {
	if c.OS == "windows" {
		p.User.Username = ec.User
	}
	return nil
}
