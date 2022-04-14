package output

import "mmesh.dev/m-api-go/grpc/resources/network"

type Interface interface {
	List(networks map[string]*network.Network)
	Show(n *network.Network)
}
type API struct{}
