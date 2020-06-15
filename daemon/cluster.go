package daemon // import "github.com/helmutkemper/moby/daemon"

import (
	lncluster "github.com/docker/libnetwork/cluster"
	apitypes "github.com/helmutkemper/moby/api/types"
	"github.com/helmutkemper/moby/api/types/filters"
)

// Cluster is the interface for github.com/helmutkemper/moby/daemon/cluster.(*Cluster).
type Cluster interface {
	ClusterStatus
	NetworkManager
	SendClusterEvent(event lncluster.ConfigEventType)
}

// ClusterStatus interface provides information about the Swarm status of the Cluster
type ClusterStatus interface {
	IsAgent() bool
	IsManager() bool
}

// NetworkManager provides methods to manage networks
type NetworkManager interface {
	GetNetwork(input string) (apitypes.NetworkResource, error)
	GetNetworks(filters.Args) ([]apitypes.NetworkResource, error)
	RemoveNetwork(input string) error
}
