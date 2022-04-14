package auth

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) PasswordReset() {
	// nxc, grpcConn := grpc.GetCoreAPIClient()
	nxc, grpcConn := grpc.GetManagerProviderAPIClient(false)
	defer grpcConn.Close()

	ur := &iam.UserRequest{
		AccountID: input.GetInput("Account ID:", "", viper.GetString("account.id"), survey.Required),
		Email:     input.GetInput("User Email:", "", viper.GetString("user.email"), input.ValidEmail),
	}

	s := output.Spinner()

	if _, err := nxc.PasswordReset(context.TODO(), ur); err != nil {
		s.Stop()
		status.Error(err, "Unable to request the password-reset")
	}

	s.Stop()

	// output.Show(sr)

	msg.Infof(`A confirmation email has been sent to %s.

This process will reset your current password with a new one auto-generated.
Your 2FA settings will be also resetted.

Follow the instructions you will find in the email and please remember to
use '%s' once you get access to change the auto-generated
password.

If for any reason you don't get the confirmation email, you can execute
this command again and the password-reset process will be restarted.

Have a nice day!`,
		colors.White(ur.Email),
		colors.White("mmeshctl user set-password"))
}
