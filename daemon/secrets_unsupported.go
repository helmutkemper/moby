// +build !linux,!windows

package daemon // import "github.com/helmutkemper/moby/daemon"

func secretsSupported() bool {
	return false
}
