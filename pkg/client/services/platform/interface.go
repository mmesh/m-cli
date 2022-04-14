package platform

import (
	"mmesh.dev/m-cli/pkg/client/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/client/services/platform/crm"
)

type Interface interface {
	Cloud() cloud.Interface
	CRM() crm.Interface
}
type API struct{}

func (a *API) Cloud() cloud.Interface {
	return &cloud.API{}
}

func (a *API) CRM() crm.Interface {
	return &crm.API{}
}
