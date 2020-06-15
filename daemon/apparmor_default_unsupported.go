// +build !linux

package daemon // import "github.com/helmutkemper/moby/daemon"

func ensureDefaultAppArmorProfile() error {
	return nil
}
