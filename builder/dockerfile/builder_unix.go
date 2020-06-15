// +build !windows

package dockerfile // import "github.com/helmutkemper/moby/builder/dockerfile"

func defaultShellForOS(os string) []string {
	return []string{"/bin/sh", "-c"}
}
