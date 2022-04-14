package tenant

import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	t := GetTenant(false)

	output.ConfirmDeletion()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	sr, err := nxc.DeleteTenant(context.TODO(), t)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete tenant")
	}

	s.Stop()

	output.Show(sr)
}
