// +build !linux,!windows

package daemon // import "github.com/helmutkemper/moby/daemon"

func configsSupported() bool {
	return false
}
