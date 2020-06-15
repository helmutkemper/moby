// +build linux

package btrfs // import "github.com/helmutkemper/moby/daemon/graphdriver/btrfs"

import (
	"testing"
)

func TestLibVersion(t *testing.T) {
	if btrfsLibVersion() <= 0 {
		t.Error("expected output from btrfs lib version > 0")
	}
}
