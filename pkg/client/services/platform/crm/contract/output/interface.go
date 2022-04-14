package output

import "mmesh.dev/m-api-go/grpc/resources/services/platform/crm"

type Interface interface {
	List(contracts map[string]*crm.Contract)
	Show(c *crm.Contract)
}
type API struct{}
