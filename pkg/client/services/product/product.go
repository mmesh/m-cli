package product

import (
	"context"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-api-go/grpc/resources/services"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/client/provider"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func GetProduct(publicCatalog bool) *services.Product {
	var pl map[string]*services.Product

	if publicCatalog {
		pl = products(nil)
	} else {
		p := &services.Provider{
			ServiceID:  viper.GetString("serviceID"),
			ProviderID: account.GetAccount().AccountID,
		}
		pl = products(p)
	}

	if len(pl) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var pOptID string
	var productOpts []string
	products := make(map[string]*services.Product)
	for _, p := range pl {
		pOptID = fmt.Sprintf("%s", p.Description)
		productOpts = append(productOpts, pOptID)
		products[pOptID] = p
	}

	pOptID = input.GetSelect("Service/Product:", "", productOpts, survey.Required)

	vars.ProductID = products[pOptID].ProductID

	return products[pOptID]
}

func products(p *services.Provider) map[string]*services.Product {
	if p == nil {
		p = provider.GetProvider(services.ProviderType_PROFESSIONAL_SERVICES)
	}

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	lr := &services.ListProductsRequest{
		Meta:     &resource.ListRequest{},
		Provider: p,
	}

	products := make(map[string]*services.Product)

	for {
		pl, err := nxc.ListProducts(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list products")
		}

		for _, p := range pl.Products {
			products[p.ProductID] = p
		}

		if len(pl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = pl.Meta.NextPageToken
		} else {
			break
		}
	}

	return products
}
