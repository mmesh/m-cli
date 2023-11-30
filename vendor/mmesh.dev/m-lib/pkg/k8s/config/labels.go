package config

import "mmesh.dev/m-api-go/grpc/resources/topology"

type AppLabel string

const (
	AppLabelController AppLabel = "mmesh-controller"
	AppLabelNode       AppLabel = "mmesh-node"
)

func (a AppLabel) String() string {
	return string(a)
}

func NodeLabels(ni *topology.NodeInstance) map[string]string {
	return map[string]string{
		"mmesh-app":       AppLabelNode.String(),
		"mmesh-type":      ni.Node.Type.String(),
		"mmesh-account":   ni.Node.AccountID,
		"mmesh-tenant":    ni.Node.TenantID,
		"mmesh-nodegroup": ni.Node.NodeGroupID,
		"mmesh-network":   ni.Node.Cfg.NetID,
		"mmesh-subnet":    ni.Node.Cfg.SubnetID,
		// "mmesh-node":      ni.Node.NodeID,
	}
}
