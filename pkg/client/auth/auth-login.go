package auth

import (
	"context"
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-cli/pkg/config"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Login(req *auth.LoginRequest, verbose bool) {
	nxc, grpcConn := grpc.GetControllerProviderAPIClient(false)
	defer grpcConn.Close()

	// if verbose {
	// 	output.ShowWaiting()
	// }

	var err error

	resp := &auth.LoginResponse{}

	ok := false
	for !ok {
		resp, err = nxc.Login(context.TODO(), req)
		if err != nil {
			status.Error(err, "Unable to login")
		}

		switch resp.Result {
		case auth.LoginResult_LOGIN_SUCCESSFUL:
			ok = true
		case auth.LoginResult_LOGIN_FAILED:
			status.Error(err, "Login failed")
		case auth.LoginResult_LOGIN_2FA_REQUIRED:
			req.TotpCode = input.GetInput("TOTP Code:", "Required only if 2FA is enabled", "", nil)
		case auth.LoginResult_IAM_ACCOUNT_UNCONFIRMED:
			grpcConn.Close()
			unconfirmedAccount()
			os.Exit(1)
		case auth.LoginResult_IAM_ACCOUNT_DISABLED:
			status.Error(err, "Account disabled, please contact customer service")
		case auth.LoginResult_IAM_USER_UNCONFIRMED:
			grpcConn.Close()
			unconfirmedUser()
			os.Exit(1)
		case auth.LoginResult_IAM_USER_DISABLED:
			status.Error(err, "User disabled, please contact your mmesh account administrator")
		}
	}

	if req.AuthMethod == auth.AuthMethod_SSH_KEY {
		resp.AuthKey.Key, err = rsaDecrypt(resp.AuthKey.Key)
		if err != nil {
			status.Error(err, "SSH authentication failed")
		}
	}

	if err := setApiKey(resp.AuthKey); err != nil {
		status.Error(err, "Unable to set apiKey")
	}

	// fmt.Println()
	if err := config.DefaultAccount(req.Realm, req.Email); err != nil {
		status.Error(err, "Unable to write configuration")
	}

	viper.Set("logged.realm", req.Realm)
	viper.Set("logged.email", req.Email)
	viper.Set("logged.isAdmin", resp.IsAdmin)

	if verbose {
		output.Authenticated()
	}
}
