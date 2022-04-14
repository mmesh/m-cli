package user

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	u := GetUser(false)

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	ur := &iam.UserRequest{
		AccountID: u.AccountID,
		Email:     u.Email,
	}

	s := output.Spinner()

	sr, err := nxc.DeleteUser(context.TODO(), ur)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete user")
	}

	s.Stop()

	output.Show(sr)
}
