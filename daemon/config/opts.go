package config // import "github.com/helmutkemper/moby/daemon/config"

import (
	"github.com/docker/swarmkit/api/genericresource"
	"github.com/helmutkemper/moby/api/types/swarm"
	"github.com/helmutkemper/moby/daemon/cluster/convert"
)

// ParseGenericResources parses and validates the specified string as a list of GenericResource
func ParseGenericResources(value []string) ([]swarm.GenericResource, error) {
	if len(value) == 0 {
		return nil, nil
	}

	resources, err := genericresource.Parse(value)
	if err != nil {
		return nil, err
	}

	obj := convert.GenericResourcesFromGRPC(resources)
	return obj, nil
}
