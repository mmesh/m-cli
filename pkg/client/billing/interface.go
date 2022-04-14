package billing

import (
	"mmesh.dev/m-cli/pkg/client/billing/invoice"
	"mmesh.dev/m-cli/pkg/client/billing/item"
)

type Interface interface {
	Item() item.Interface
	Invoice() invoice.Interface
}
type API struct{}

func (a *API) Item() item.Interface {
	return &item.API{}
}

func (a *API) Invoice() invoice.Interface {
	return &invoice.API{}
}
