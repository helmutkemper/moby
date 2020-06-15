// +build !linux

package daemon // import "github.com/helmutkemper/moby/daemon"

func selinuxSetDisabled() {
}

func selinuxFreeLxcContexts(label string) {
}

func selinuxEnabled() bool {
	return false
}
