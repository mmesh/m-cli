package cloud

import (
	"mmesh.dev/m-cli/pkg/client/services/platform/cloud/instance"
	"mmesh.dev/m-cli/pkg/client/services/platform/cloud/k8s"
)

type Interface interface {
	Instance() instance.Interface
	Kubernetes() k8s.Interface
}
type API struct{}

func (a *API) Instance() instance.Interface {
	return &instance.API{}
}

func (a *API) Kubernetes() k8s.Interface {
	return &k8s.API{}
}
