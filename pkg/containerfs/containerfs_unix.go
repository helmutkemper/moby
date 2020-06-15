// +build !windows

package containerfs // import "github.com/helmutkemper/moby/pkg/containerfs"

import "path/filepath"

// cleanScopedPath preappends a to combine with a mnt path.
func cleanScopedPath(path string) string {
	return filepath.Join(string(filepath.Separator), path)
}
