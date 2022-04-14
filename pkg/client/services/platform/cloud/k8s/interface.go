package k8s

import "mmesh.dev/m-cli/pkg/client/services/platform/cloud/k8s/cluster"

type Interface interface {
	Cluster() cluster.Interface
}
type API struct{}

func (a *API) Cluster() cluster.Interface {
	return &cluster.API{}
}
