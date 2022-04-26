package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/internal/app/cli/auth/login"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// userShowCmd represents the user get verb
var userShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show user details",
	Long:  appHeader(`Show user details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.User().ShowLoggedUser(login.NewRequest())
	},
}

// userSetEmailCmd represents the user set-password verb
var userSetEmailCmd = &cobra.Command{
	Use:   "set-email",
	Short: "Set user email",
	Long:  appHeader(`Set user email.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.User().SetEmail(login.NewRequest())
	},
}

// userSetPasswordCmd represents the user set-password verb
var userSetPasswordCmd = &cobra.Command{
	Use:   "set-password",
	Short: "Set user password",
	Long:  appHeader(`Set user password.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.User().SetPassword(login.NewRequest())
	},
}

// userSetSSHKeysCmd represents the user set-ssh-keys verb
var userSetSSHKeysCmd = &cobra.Command{
	Use:   "set-ssh-keys",
	Short: "Set user SSH keys",
	Long:  appHeader(`Set user SSH keys.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.User().SetSSHKeys(login.NewRequest())
	},
}

// userSetTOTPCmd represents the iam/users set-totp verb
var userSetTOTPCmd = &cobra.Command{
	Use:   "set-2fa",
	Short: "Set user 2-factor authentication",
	Long:  appHeader(`Set user 2-factor authentication.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.User().SetTOTP(login.NewRequest())
	},
}

// userPasswordResetCmd represents the password-reset command
var userPasswordResetCmd = &cobra.Command{
	Use:   "password-reset",
	Short: "Request a password-reset",
	Long:  appHeader(`Request a password-reset.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Auth().PasswordReset()
	},
}

func init() {
	userCmd.AddCommand(userShowCmd)
	userCmd.AddCommand(userSetEmailCmd)
	userCmd.AddCommand(userSetPasswordCmd)
	userCmd.AddCommand(userSetSSHKeysCmd)
	userCmd.AddCommand(userSetTOTPCmd)
	userCmd.AddCommand(userPasswordResetCmd)

	userCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
}
