package auth

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) Login(req *auth.LoginRequest, verbose bool) {
	nxc, grpcConn := grpc.GetManagerProviderAPIClient(false)
	// defer grpcConn.Close()

	var resp *auth.LoginResponse
	var err error

	for {
		resp, err = nxc.Login(context.TODO(), req)
		if err != nil {
			grpcConn.Close()
			status.Error(err, "Unable to login")
		}

		if resp.Result == auth.LoginResult_LOGIN_2FA_REQUIRED {
			req.TotpCode = input.GetInput("TOTP Code:", "Required only if 2FA is enabled", "", nil)
		} else {
			grpcConn.Close()
			break
		}
	}

	switch resp.Result {
	case auth.LoginResult_LOGIN_FAILED:
		msg.Error("Login failed")
	case auth.LoginResult_IAM_ACCOUNT_UNCONFIRMED:
		unconfirmedAccount()
	case auth.LoginResult_IAM_ACCOUNT_DISABLED:
		msg.Error("Account disabled, please contact customer service")
	case auth.LoginResult_IAM_USER_UNCONFIRMED:
		unconfirmedUser()
	case auth.LoginResult_IAM_USER_DISABLED:
		msg.Error("User disabled, please contact your mmesh account administrator")
	}

	if resp.Result != auth.LoginResult_LOGIN_SUCCESSFUL {
		os.Exit(1)
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

	viper.Set("logged.realm", req.Realm)
	viper.Set("logged.email", req.Email)
	viper.Set("logged.isAdmin", resp.IsAdmin)

	if verbose {
		fmt.Println()
		output.Authenticated()
	}
}
