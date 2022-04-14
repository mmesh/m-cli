package sg

import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	sg := GetSecurityGroup(false)

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	sr, err := nxc.DeleteSecurityGroup(context.TODO(), sg)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete security-group")
	}

	s.Stop()

	output.Show(sr)
}
