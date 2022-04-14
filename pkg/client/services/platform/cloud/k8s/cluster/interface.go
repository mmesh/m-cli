package cluster

import "mmesh.dev/m-cli/pkg/client/services/platform/cloud/k8s/cluster/output"

type Interface interface {
	List()
	Show()
	Create()
	Delete()

	GetKubeConfig()
	CreateNodePool()
	DeleteNodePool()
	AddNode()
	DeleteNode()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
