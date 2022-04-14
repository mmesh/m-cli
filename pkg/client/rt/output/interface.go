package output

import "mmesh.dev/m-api-go/grpc/network/mmnp/routing"

type Interface interface {
	List(rt *routing.RoutingTable)
}
type API struct{}
