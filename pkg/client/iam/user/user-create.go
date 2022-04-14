package user

import (
	"context"
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/client/iam/user/credentials"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) Create() {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	u := &iam.User{
		AccountID: a.AccountID,
		Email:     input.GetInput("Email:", "", "", input.ValidEmail),
		Username:  input.GetInput("Username:", "", "", input.ValidID),
		Credentials: &iam.UserCredentials{
			Password: credentials.SetCredentialsPassword(true),
			SSH:      &iam.SSH{},
			TOTP:     &iam.TOTP{},
		},
		RBAC:   &iam.UserRBAC{},
		Status: &iam.UserStatus{},
	}

	u.RBAC = setUserRBAC(nxc, u)

	s := output.Spinner()

	u, err := nxc.CreateUser(context.TODO(), u)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to create user")
	}

	s.Stop()

	Output().Show(u)

	msg.Infof("User %v created :-)", colors.White(u.Email))

	fmt.Println()
}
