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

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	output.ConfirmDeletion()

	ur := &iam.UserReq{
		AccountID: u.AccountID,
		UserID:    u.UserID,
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
