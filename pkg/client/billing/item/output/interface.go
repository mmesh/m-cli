package output

import "mmesh.dev/m-api-go/grpc/resources/billing"

type Interface interface {
	List(invoiceItems map[string]*billing.InvoiceItem)
	Show(bi *billing.InvoiceItem)
}
type API struct{}
