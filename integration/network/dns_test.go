package network // import "github.com/helmutkemper/moby/integration/network"

import (
	"context"
	"testing"
	"time"

	"github.com/helmutkemper/moby/api/types"
	"github.com/helmutkemper/moby/integration/internal/container"
	"github.com/helmutkemper/moby/integration/internal/network"
	"github.com/helmutkemper/moby/testutil/daemon"
	"gotest.tools/v3/poll"
	"gotest.tools/v3/skip"
)

func TestDaemonDNSFallback(t *testing.T) {
	skip.If(t, testEnv.IsRemoteDaemon, "cannot start daemon on remote test run")
	skip.If(t, testEnv.DaemonInfo.OSType != "linux")
	skip.If(t, IsUserNamespace())

	d := daemon.New(t)
	d.StartWithBusybox(t, "-b", "none", "--dns", "127.127.127.1", "--dns", "8.8.8.8")
	defer d.Stop(t)

	c := d.NewClientT(t)
	ctx := context.Background()

	network.CreateNoError(ctx, t, c, "test")
	defer c.NetworkRemove(ctx, "test")

	cid := container.Run(ctx, t, c, container.WithNetworkMode("test"), container.WithCmd("nslookup", "docker.com"))
	defer c.ContainerRemove(ctx, cid, types.ContainerRemoveOptions{Force: true})

	poll.WaitOn(t, container.IsSuccessful(ctx, c, cid), poll.WithDelay(100*time.Millisecond), poll.WithTimeout(2*time.Second))
}
