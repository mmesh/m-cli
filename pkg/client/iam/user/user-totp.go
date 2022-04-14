package user

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	auth_pb "mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-cli/pkg/client/auth"
	"mmesh.dev/m-cli/pkg/client/iam/user/credentials"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) SetTOTP(loginReq *auth_pb.LoginRequest) {
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

	suctr := &iam.SetUserCredentialsTOTPRequest{
		AccountID:       u.AccountID,
		Email:           u.Email,
		TOTPCredentials: credentials.SetCredentialsTOTP(u),
	}

	s := output.Spinner()

	u, err = nxc.SetUserCredentialsTOTP(context.TODO(), suctr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set user 2FA credentials")
	}

	s.Stop()

	Output().Show(u)

	msg.Infof("User %v configured :-)", colors.White(u.Email))
}
