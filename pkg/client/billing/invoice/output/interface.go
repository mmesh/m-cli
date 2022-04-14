package output

import "mmesh.dev/m-api-go/grpc/resources/billing"

type Interface interface {
	List(invoices map[string]*billing.Invoice)
	Show(i *billing.Invoice)
}
type API struct{}
