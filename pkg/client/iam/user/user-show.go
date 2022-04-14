package user

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	auth_pb "mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-cli/pkg/client/auth"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Show() {
	Output().Show(GetUser(false))
}

func (api *API) ShowLoggedUser(loginReq *auth_pb.LoginRequest) {
	auth.Resource().Login(loginReq, false)

	lu := getLoggedUser()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	ur := &iam.UserRequest{
		AccountID: lu.accountID,
		Email:     lu.email,
	}

	u, err := nxc.GetUser(context.TODO(), ur)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get user")
	}

	s.Stop()

	Output().Show(u)
}
