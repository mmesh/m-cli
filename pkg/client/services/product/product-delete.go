package product

import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	p := GetProduct(false)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	sr, err := nxc.DeleteProduct(context.TODO(), p)
	if err != nil {
		status.Error(err, "Unable to delete product")
	}

	s.Stop()

	output.Show(sr)
}
