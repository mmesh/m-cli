package user

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	auth_pb "mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-cli/pkg/client/auth"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) SetEmail(loginReq *auth_pb.LoginRequest) {
	auth.Resource().Login(loginReq, false)

	lu := getLoggedUser()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	ur := &iam.UserRequest{
		AccountID: lu.accountID,
		Email:     lu.email,
	}

	u, err := nxc.GetUser(context.TODO(), ur)
	if err != nil {
		status.Error(err, "Unable to get user")
	}

	suer := &iam.SetUserEmailRequest{
		AccountID: u.AccountID,
		Email:     u.Email,
		NewEmail:  input.GetInput("Email:", "", u.Email, input.ValidEmail),
	}

	if suer.NewEmail == u.Email {
		return
	}

	s := output.Spinner()

	u, err = nxc.SetUserEmail(context.TODO(), suer)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set user email")
	}

	s.Stop()

	Output().Show(u)

	msg.Infof("User %v configured :-)", colors.White(u.Email))
}
