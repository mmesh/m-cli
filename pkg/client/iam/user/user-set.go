package user

/*
import (
	"context"

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

func (api *API) Set() {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	u := &iam.User{
		AccountID: a.AccountID,
		Credentials: &iam.UserCredentials{
			SSH:  &iam.SSH{},
			TOTP: &iam.TOTP{},
		},
	}

	newUser := false

	usr := GetUser(true)
	if usr != nil { // editing existing resource
		output.Choice("Edit User")

		u = usr
	} else { // <new> resource
		output.Choice("New User")

		u.AccountID = a.AccountID

		newUser = true
	}

	u.NewEmail = input.GetInput("Email:", "", u.Email, input.ValidEmail)

	if len(u.Email) == 0 {
		u.Email = u.NewEmail
		u.NewEmail = ""
	} else {
		if u.Email == u.NewEmail {
			u.NewEmail = ""
		}
	}

	u.Username = input.GetInput("Username:", "", u.Username, input.ValidID)

	if u.Credentials == nil {
		u.Credentials = &iam.UserCredentials{}
	}
	pw := credentials.SetCredentialsPassword(newUser)

	if len(pw) > 0 {
		u.Credentials.Password = pw
	}

	u = credentials.SetSSHKeys(u, true)

	if u.Credentials.TOTP == nil {
		u.Credentials.TOTP = &iam.TOTP{}
	}
	u.Credentials.TOTP = credentials.SetCredentialsTOTP(u)

	if u.Status == nil {
		u.Status = &iam.UserStatus{}
	}

	if newUser {
		u.RBAC = setUserRBAC(nxc, u)
	}

	s := output.Spinner()

	if _, err := nxc.SetUser(context.TODO(), u); err != nil {
		status.Error(err, "Unable to set user")
	}

	if len(pw) > 0 {
		if _, err := nxc.SetUserCredentialsPassword(context.TODO(), u); err != nil {
			status.Error(err, "Unable to set user credentials")
		}
	}

	u, err := nxc.SetUserCredentialsSSH(context.TODO(), u)
	if err != nil {
		status.Error(err, "Unable to set user SSH credentials")
	}

	u, err = nxc.SetUserSSHKeys(context.TODO(), u)
	if err != nil {
		status.Error(err, "Unable to set user SSH keys")
	}

	u, err = nxc.SetUserCredentialsTOTP(context.TODO(), u)
	if err != nil {
		status.Error(err, "Unable to set user 2FA credentials")
	}

	if newUser {
		u, err = nxc.SetUserPermissions(context.TODO(), u)
		if err != nil {
			status.Error(err, "Unable to set user permissions")
		}
	}

	s.Stop()

	Output().Show(u)

	msg.Infof("User %v configured :-)", colors.White(u.Email))
}
*/
