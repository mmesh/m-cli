package config

import (
	"mmesh.dev/m-api-go/grpc/resources/topology"
)

type AppLabel string

const (
	AppLabelController AppLabel = "mmesh-controller"
	AppLabelNode       AppLabel = "mmesh-node"
)

func (a AppLabel) String() string {
	return string(a)
}

type InstanceLabelType string

const (
	InstanceLabelTypeGeneric           InstanceLabelType = "generic"
	InstanceLabelTypeKubernetesGateway InstanceLabelType = "k8sgw"
)

func (a InstanceLabelType) String() string {
	return string(a)
}

func GetInstanceLabelType(t topology.NodeType) string {
	switch t {
	case topology.NodeType_GENERIC:
		return InstanceLabelTypeGeneric.String()
	case topology.NodeType_K8S_GATEWAY:
		return InstanceLabelTypeKubernetesGateway.String()
	// case topology.NodeType_CLOUD_INSTANCE:
	// 	return InstanceLabelTypeGeneric.String()
	}

	return InstanceLabelTypeGeneric.String()
}
