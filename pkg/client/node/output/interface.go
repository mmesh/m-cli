package output

import "mmesh.dev/m-api-go/grpc/resources/topology"

type Interface interface {
	List(nodes map[string]*topology.Node)
	Show(n *topology.Node)
	Metrics(n *topology.Node)
	ShowEndpoint(e *topology.Endpoint)
}
type API struct{}
