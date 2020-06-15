package daemon // import "github.com/helmutkemper/moby/daemon"

import (
	"context"

	"github.com/helmutkemper/moby/api/types"
	"github.com/helmutkemper/moby/dockerversion"
)

// AuthenticateToRegistry checks the validity of credentials in authConfig
func (daemon *Daemon) AuthenticateToRegistry(ctx context.Context, authConfig *types.AuthConfig) (string, string, error) {
	return daemon.RegistryService.Auth(ctx, authConfig, dockerversion.DockerUserAgent(ctx))
}
