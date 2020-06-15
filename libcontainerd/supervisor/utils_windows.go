package supervisor // import "github.com/helmutkemper/moby/libcontainerd/supervisor"

import "syscall"

// containerdSysProcAttr returns the SysProcAttr to use when exec'ing
// containerd
func containerdSysProcAttr() *syscall.SysProcAttr {
	return nil
}
