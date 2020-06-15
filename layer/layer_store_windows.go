package layer // import "github.com/helmutkemper/moby/layer"

import (
	"io"

	"github.com/docker/distribution"
)

func (ls *layerStore) RegisterWithDescriptor(ts io.Reader, parent ChainID, descriptor distribution.Descriptor) (Layer, error) {
	return ls.registerWithDescriptor(ts, parent, descriptor)
}
