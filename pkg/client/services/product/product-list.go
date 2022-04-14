package product

import (
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/services"
	"mmesh.dev/m-cli/pkg/client/account"
)

func (api *API) List(publicCatalog bool) {
	if publicCatalog {
		Output().List(products(nil))
	} else {
		p := &services.Provider{
			ServiceID:  viper.GetString("serviceID"),
			ProviderID: account.GetAccount().AccountID,
		}
		Output().List(products(p))
	}
}
