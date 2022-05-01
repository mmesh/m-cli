package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/internal/app/cli/auth/login"
	"mmesh.dev/m-cli/pkg/client"
)

// authCmd represents the auth command
var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authentication commands",
	Long:  appHeader(`User session authentication commands.`),
}

// authLoginCmd represents the login command
var authLoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Log in",
	Long:  appHeader(`Log in.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Auth().Login(login.NewRequest(), true)
	},
}

// authLogoutCmd represents the logout command
var authLogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Log out",
	Long:  appHeader(`Log out.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Auth().Logout()
	},
}

// authPasswordResetCmd represents the password-reset command
var authPasswordResetCmd = &cobra.Command{
	Use:   "password-reset",
	Short: "Request a password-reset",
	Long:  appHeader(`Request a password-reset.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		header()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Auth().PasswordReset()
	},
}

// authConfirmationMailResendCmd represents the password-reset command
var authConfirmationMailResendCmd = &cobra.Command{
	Use:   "resend-confirmation",
	Short: "Request a new confirmation mail",
	Long:  appHeader(`Request a new confirmation mail.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		header()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Auth().ConfirmationMailResend()
	},
}

/*
// authTokenCmd represents the login command
var authTokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Generate an authorization token",
	Long:  appHeader(`Generate an authorization token.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Auth().Token()
	},
}
*/

func init() {
	authCmd.AddCommand(authLoginCmd)
	authCmd.AddCommand(authLogoutCmd)
	authCmd.AddCommand(authPasswordResetCmd)
	authCmd.AddCommand(authConfirmationMailResendCmd)
	// authCmd.AddCommand(authTokenCmd)
}
