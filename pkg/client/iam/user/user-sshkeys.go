package user

import (
	"context"
	"fmt"

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

func (api *API) SetSSHKeys(loginReq *auth_pb.LoginRequest) {
	auth.Resource().Login(loginReq, true)

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

	u = credentials.SetSSHKeys(u, true)

	s := output.Spinner()

	if u.Credentials.SSH == nil {
		if _, err := nxc.DeleteUserCredentialsSSH(context.TODO(), ur); err != nil {
			s.Stop()
			status.Error(err, "Unable to delete user SSH credentials")
		}
	} else {
		sucsr := &iam.SetUserCredentialsSSHRequest{
			AccountID:      u.AccountID,
			Email:          u.Email,
			SSHCredentials: u.Credentials.SSH,
		}

		if _, err := nxc.SetUserCredentialsSSH(context.TODO(), sucsr); err != nil {
			s.Stop()
			status.Error(err, "Unable to set user SSH credentials")
		}
	}

	suskr := &iam.SetUserSSHKeysRequest{
		AccountID: u.AccountID,
		Email:     u.Email,
		SSHKeys:   u.SSHKeys,
	}

	u, err = nxc.SetUserSSHKeys(context.TODO(), suskr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set user SSH keys")
	}

	s.Stop()

	Output().Show(u)

	msg.Infof("SSH keys configured for user %v :-)", colors.DarkWhite(u.Email))

	if u.Credentials.SSH != nil {
		if u.Credentials.SSH.Enabled && u.Credentials.SSH.Key != nil {
			sshAuthHelp(u.AccountID)
		}
	}
}

func (api *API) GetSSHKey(loginReq *auth_pb.LoginRequest) *iam.SSHKey {
	auth.Resource().Login(loginReq, true)

	lu := getLoggedUser()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	ur := &iam.UserRequest{
		AccountID: lu.accountID,
		Email:     lu.email,
	}

	userSSHKeys, err := nxc.GetUserSSHKeys(context.TODO(), ur)
	if err != nil {
		status.Error(err, "Unable to get user SSH keys")
	}

	return credentials.GetSSHKey(userSSHKeys)
}

func sshAuthHelp(accountID string) {
	sshKeyFile := fmt.Sprintf("$HOME/.mmesh/%s/id_rsa", accountID)
	q := colors.DarkBlue("'")

	fmt.Printf(`Please, remember that mmeshctl will look for the private key
for authentication at %s%s%s.

For security, it is highly recommended to limit the r/w permissions
of that file (%s%s%s).

`,
		q, colors.DarkWhite(sshKeyFile), q,
		q, colors.DarkWhite(fmt.Sprintf("chmod 0400 %s", sshKeyFile)), q,
	)
}
