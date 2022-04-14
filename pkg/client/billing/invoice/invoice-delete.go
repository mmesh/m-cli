package invoice

import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	i := GetInvoice()

	nxc, grpcConn := grpc.GetBillingAPIClient(true)
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	sr, err := nxc.DeleteInvoice(context.TODO(), i)
	if err != nil {
		status.Error(err, "Unable to delete invoice")
	}

	s.Stop()

	output.Show(sr)
}
