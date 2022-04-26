package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// iamUsersCmd represents the iamUsers command
var iamUsersCmd = &cobra.Command{
	Use:   "user",
	Short: "IAM users adminstration",
	Long:  appHeader(`IAM users adminstration operations.`),
}

// iamUsersListCmd represents the iam/users list verb
var iamUsersListCmd = &cobra.Command{
	Use:   "list",
	Short: "List users",
	Long:  appHeader(`List all realm's users.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.User().List()
	},
}

// iamUsersShowCmd represents the iam/users get verb
var iamUsersShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show user",
	Long:  appHeader(`Show user details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.User().Show()
	},
}

// iamUsersCreateCmd represents the iam/users set verb
var iamUsersCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a user",
	Long:  appHeader(`Create a user.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.User().Create()
	},
}

// iamUsersDeleteCmd represents the iam/users delete verb
var iamUsersDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete user",
	Long:  appHeader(`Remove user from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.User().Delete()
	},
}

// iamUsersSetPermissionsCmd represents the iam/users set verb
var iamUsersSetPermissionsCmd = &cobra.Command{
	Use:   "set-perms",
	Short: "Set user RBAC permissions (admin only)",
	Long:  appHeader(`Set user RBAC permissions (admin only).`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.User().SetPermissions()
	},
}

// iamUsersEnableCmd represents the iam/user enable verb
var iamUsersEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable user",
	Long:  appHeader(`Enable user.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.User().Enable()
	},
}

// iamUsersDisableCmd represents the iam/user disable verb
var iamUsersDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable user",
	Long:  appHeader(`Disable user.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.User().Disable()
	},
}

func init() {
	iamCmd.AddCommand(iamUsersCmd)
	iamUsersCmd.AddCommand(iamUsersListCmd)
	iamUsersCmd.AddCommand(iamUsersShowCmd)
	iamUsersCmd.AddCommand(iamUsersCreateCmd)
	iamUsersCmd.AddCommand(iamUsersDeleteCmd)
	iamUsersCmd.AddCommand(iamUsersEnableCmd)
	iamUsersCmd.AddCommand(iamUsersDisableCmd)
	iamUsersCmd.AddCommand(iamUsersSetPermissionsCmd)

	iamUsersCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
}
