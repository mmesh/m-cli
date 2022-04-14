package node

import (
	"mmesh.dev/m-api-go/grpc/resources/network"
)

type Interface interface {
	CreateGateway(ni *network.NodeInstance) error
	DeleteGateway(ns, name string) error

	ConnectStatefulSet(ns, name string, ni *network.NodeInstance) error
	DisconnectStatefulSet(ns, name string, ni *network.NodeInstance) error
	ConnectDeployment(ns, name string, ni *network.NodeInstance) error
	DisconnectDeployment(ns, name string, ni *network.NodeInstance) error
	ConnectDaemonSet(ns, name string, ni *network.NodeInstance) error
	DisconnectDaemonSet(ns, name string, ni *network.NodeInstance) error
}

type API struct {
	KubeConfig []byte
}
