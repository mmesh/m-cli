package auth

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) ConfirmationMailResend() {
	nxc, grpcConn := grpc.GetManagerProviderAPIClient(false)
	defer grpcConn.Close()

	ur := &iam.UserRequest{
		AccountID: input.GetInput("Account ID:", "", viper.GetString("account.id"), survey.Required),
		Email:     input.GetInput("User Email:", "", viper.GetString("user.email"), input.ValidEmail),
	}

	s := output.Spinner()

	if _, err := nxc.ConfirmationMailResend(context.TODO(), ur); err != nil {
		s.Stop()
		status.Error(err, "Unable to request a new confirmation mail")
	}

	s.Stop()

	// output.Show(sr)

	fmt.Printf(`
A confirmation email has been sent to %s.

Please follow the instructions you will find in the email to confirm
your account.

If for any reason you don't get the confirmation email, you can
execute this command again and a new email will be send.

Have a nice day!

`,
		colors.White(ur.Email))
}

func unconfirmedAccount() {
	fmt.Printf(`
Your mmesh account is not confirmed, if you are the account
admin please check your inbox to find the confirmation email and
follow the instructions to activate your account.

If for any reason you didn't get the confirmation email, you can
execute the command '%s' and a
new email will be send to the account admin.

Have a nice day!

`,
		colors.White("mmeshctl auth resend-confirmation"))
}

func unconfirmedUser() {
	fmt.Printf(`
Your mmesh user email is not confirmed, please check your
inbox to find the confirmation email and follow the instructions
to activate your account.

If for any reason you didn't get the confirmation email, you can
execute the command '%s' and a
new email will be send.

Have a nice day!

`,
		colors.White("mmeshctl auth resend-confirmation"))
}
