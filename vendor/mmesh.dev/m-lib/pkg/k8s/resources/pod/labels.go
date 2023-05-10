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

	return map[string]string{
		// "mmesh-federation": ni.FederationID,
		string(appLabel):   ni.Node.NodeID,
		"mmesh-app":        string(appLabel),
		"mmesh-objectID":   ni.Node.NodeID,
		"mmesh-type":       config.GetInstanceLabelType(ni.Type),
		"mmesh-account":    ni.Node.AccountID,
		"mmesh-tenant":     ni.Node.TenantID,
		"mmesh-network":    ni.Node.NetID,
		"mmesh-subnet":     ni.Node.SubnetID,
	}
}
