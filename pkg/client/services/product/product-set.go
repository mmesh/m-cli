package product

import (
	"context"
	"errors"
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/services/catalog/pro"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils"
)

func (api *API) Set(yamlFile string) {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	pc := &pro.ProductCatalog{}

	if err := utils.FileParser(yamlFile, pc); err != nil {
		status.Error(err, "Unable to parse YAML file")
	}

	if a.AccountID != pc.ProviderID {
		err := errors.New("ProviderID in file does not match your accountID")
		status.Error(err, "Invalid ProviderID")
	}

	// if err := validProductID(p.ProductID); err != nil {
	// 	status.Error(err, "Invalid ProductID")
	// }

	pc.ServiceID = "mmesh"
	pc.ProviderID = a.AccountID

	s := output.Spinner()

	_, err := nxc.SetProductCatalog(context.TODO(), pc)
	if err != nil {
		status.Error(err, "Unable to set product catalog")
	}

	s.Stop()

	fmt.Printf(`

Your catalog has been published!

Thanks for using mmesh!

`)
}
