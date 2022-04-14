package output

import "mmesh.dev/m-api-go/grpc/resources/network"

type Interface interface {
	List(vrfs map[string]*network.VRF)
	Show(v *network.VRF)
}
type API struct{}
