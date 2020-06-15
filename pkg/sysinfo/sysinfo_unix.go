// +build !linux,!windows

package sysinfo // import "github.com/helmutkemper/moby/pkg/sysinfo"

type opts struct{}

// Opt for New().
type Opt func(*opts)

// New returns an empty SysInfo for non linux for now.
func New(quiet bool, options ...Opt) *SysInfo {
	sysInfo := &SysInfo{}
	return sysInfo
}
