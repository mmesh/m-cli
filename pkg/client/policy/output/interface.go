package output

import "mmesh.dev/m-api-go/grpc/resources/topology"

type Interface interface {
	Show(s *topology.Subnet)
}
type API struct{}
