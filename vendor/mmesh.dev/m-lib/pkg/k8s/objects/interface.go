package objects

import (
	"mmesh.dev/m-lib/pkg/k8s/objects/controller"
	"mmesh.dev/m-lib/pkg/k8s/objects/node"
)

type Interface interface {
	Controller() controller.Interface
	Node() node.Interface
}

type API struct {
	KubeConfig []byte
}

func (a *API) Controller() controller.Interface {
	return &controller.API{KubeConfig: a.KubeConfig}
}

func (a *API) Node() node.Interface {
	return &node.API{KubeConfig: a.KubeConfig}
}
