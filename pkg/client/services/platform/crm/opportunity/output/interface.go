package output

import "mmesh.dev/m-api-go/grpc/resources/services/platform/crm"

type Interface interface {
	List(opportunities map[string]*crm.Opportunity)
	Show(o *crm.Opportunity)
}
type API struct{}
