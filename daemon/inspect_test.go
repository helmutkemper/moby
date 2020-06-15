package daemon // import "github.com/helmutkemper/moby/daemon"

import (
	"testing"

	containertypes "github.com/helmutkemper/moby/api/types/container"
	"github.com/helmutkemper/moby/container"
	"github.com/helmutkemper/moby/daemon/config"
	"github.com/helmutkemper/moby/daemon/exec"
	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetInspectData(t *testing.T) {
	c := &container.Container{
		ID:           "inspect-me",
		HostConfig:   &containertypes.HostConfig{},
		State:        container.NewState(),
		ExecCommands: exec.NewStore(),
	}

	d := &Daemon{
		linkIndex:   newLinkIndex(),
		configStore: &config.Config{},
	}

	_, err := d.getInspectData(c)
	assert.Check(t, is.ErrorContains(err, ""))

	c.Dead = true
	_, err = d.getInspectData(c)
	assert.Check(t, err)
}
