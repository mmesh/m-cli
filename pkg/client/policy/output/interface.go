package output

import "mmesh.dev/m-api-go/grpc/resources/network"

type Interface interface {
	Show(np *network.Policy)
}
type API struct{}
