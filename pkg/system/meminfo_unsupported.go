// +build !linux,!windows

package system // import "github.com/helmutkemper/moby/pkg/system"

// ReadMemInfo is not supported on platforms other than linux and windows.
func ReadMemInfo() (*MemInfo, error) {
	return nil, ErrNotSupportedPlatform
}
