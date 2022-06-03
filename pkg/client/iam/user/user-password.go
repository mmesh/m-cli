package user

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	auth_pb "mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-cli/pkg/client/auth"
	"mmesh.dev/m-cli/pkg/client/iam/user/credentials"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) SetPassword(loginReq *auth_pb.LoginRequest) {
	auth.Resource().Login(loginReq, true)

	lu := getLoggedUser()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	sucpr := &iam.SetUserCredentialsPasswordRequest{
		AccountID:       lu.accountID,
		Email:           lu.email,
		CurrentPassword: input.GetPassword("Current Password:", ""),
		NewPassword:     credentials.SetCredentialsPassword(false),
	}

	if len(sucpr.NewPassword) == 0 {
		return
	}

	s := output.Spinner()

	sr, err := nxc.SetUserCredentialsPassword(context.TODO(), sucpr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set user password")
	}

	s.Stop()

	output.Show(sr)
}
