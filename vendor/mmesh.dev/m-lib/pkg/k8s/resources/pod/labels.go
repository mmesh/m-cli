package pod

import (
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-lib/pkg/k8s/config"
)

func (a *API) NewLabels(i interface{}, appLabel config.AppLabel) map[string]string {
	var ni *topology.NodeInstance

	switch appLabel {
	case config.AppLabelNode:
		ni = i.(*topology.NodeInstance)
	default:
		return nil
	}

	return config.NodeLabels(ni)
}
