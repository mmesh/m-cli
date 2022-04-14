package config

import (
	"mmesh.dev/m-api-go/grpc/resources/network"
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
	InstanceLabelTypeMRS               InstanceLabelType = "mrs"
)

func (a InstanceLabelType) String() string {
	return string(a)
}

func GetInstanceLabelType(t network.NodeInstanceType) string {
	switch t {
	case network.NodeInstanceType_GENERIC:
		return InstanceLabelTypeGeneric.String()
	case network.NodeInstanceType_RELAY_NODE:
		return InstanceLabelTypeGeneric.String()
	case network.NodeInstanceType_K8S_GATEWAY:
		return InstanceLabelTypeKubernetesGateway.String()
	case network.NodeInstanceType_K8S_RELAY_SERVICE:
		return InstanceLabelTypeMRS.String()
	case network.NodeInstanceType_CLOUD_INSTANCE:
		return InstanceLabelTypeGeneric.String()
	}

	return InstanceLabelTypeGeneric.String()
}
