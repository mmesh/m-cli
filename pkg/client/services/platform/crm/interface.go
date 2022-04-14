package crm

import (
	"mmesh.dev/m-cli/pkg/client/services/platform/crm/contract"
	"mmesh.dev/m-cli/pkg/client/services/platform/crm/opportunity"
)

type Interface interface {
	Opportunity() opportunity.Interface
	Contract() contract.Interface
}
type API struct{}

func (a *API) Opportunity() opportunity.Interface {
	return &opportunity.API{}
}

func (a *API) Contract() contract.Interface {
	return &contract.API{}
}
