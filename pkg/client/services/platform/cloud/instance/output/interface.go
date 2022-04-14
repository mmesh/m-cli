package output

import "mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"

type Interface interface {
	List(instances map[string]*cloud.Instance)
	Show(i *cloud.Instance)
}
type API struct{}
