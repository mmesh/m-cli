package output

import "mmesh.dev/m-api-go/grpc/resources/network"

type Interface interface {
	Show(tenantID, netID, vrfIF string, np *network.Policy)
}
type API struct{}
