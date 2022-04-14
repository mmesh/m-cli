package output

import "mmesh.dev/m-api-go/grpc/resources/network"

type Interface interface {
	List(nodes map[string]*network.Node)
	Show(n *network.Node)
	Metrics(n *network.Node)
	ShowEndpoint(e *network.NetworkEndpoint)
}
type API struct{}
