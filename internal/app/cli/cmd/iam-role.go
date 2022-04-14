package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// iamRolesCmd represents the iamRoles command
var iamRolesCmd = &cobra.Command{
	Use:   "role",
	Short: "IAM roles adminstration",
	Long:  `IAM roles adminstration operations.`,
}

// iamRolesListCmd represents the iam/roles list verb
var iamRolesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List roles",
	Long:  `List all roles.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Role().List()
	},
}

// iamRolesShowCmd represents the iam/roles get verb
var iamRolesShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show role",
	Long:  `Show role details.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Role().Show()
	},
}

// iamRolesSetCmd represents the iam/roles set verb
var iamRolesSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Create or update role",
	Long:  `Create or update role.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Role().Set()
	},
}

// iamRolesDeleteCmd represents the iam/roles delete verb
var iamRolesDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete role",
	Long:  `Delete role.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Role().Delete()
	},
}

func init() {
	iamCmd.AddCommand(iamRolesCmd)
	iamRolesCmd.AddCommand(iamRolesListCmd)
	iamRolesCmd.AddCommand(iamRolesShowCmd)
	iamRolesCmd.AddCommand(iamRolesSetCmd)
	iamRolesCmd.AddCommand(iamRolesDeleteCmd)

	iamRolesCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
}
