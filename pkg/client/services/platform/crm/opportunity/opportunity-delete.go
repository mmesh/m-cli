package opportunity

import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	o := GetOpportunity()

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	sr, err := nxc.DeleteOpportunity(context.TODO(), o)
	if err != nil {
		status.Error(err, "Unable to delete CRM opportunity")
	}

	s.Stop()

	output.Show(sr)
}
