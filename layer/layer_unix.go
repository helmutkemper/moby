// +build linux freebsd darwin openbsd

package layer // import "github.com/helmutkemper/moby/layer"

import "github.com/helmutkemper/moby/pkg/stringid"

func (ls *layerStore) mountID(name string) string {
	return stringid.GenerateRandomID()
}
