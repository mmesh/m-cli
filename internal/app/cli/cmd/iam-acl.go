package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// iamACLsCmd represents the iamACLs command
var iamACLsCmd = &cobra.Command{
	Use:   "acl",
	Short: "IAM ACLs adminstration",
	Long:  `IAM ACLs adminstration operations.`,
}

// iamACLsListCmd represents the iam/roles list verb
var iamACLsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List ACLs",
	Long:  `List all ACLs.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.ACL().List()
	},
}

// iamACLsShowCmd represents the iam/roles get verb
var iamACLsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show ACL",
	Long:  `Show ACL details.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.ACL().Show()
	},
}

// iamACLsSetCmd represents the iam/roles set verb
var iamACLsSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Create or update ACL",
	Long:  `Create or update ACL.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.ACL().Set()
	},
}

// iamACLsDeleteCmd represents the iam/roles delete verb
var iamACLsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete ACL",
	Long:  `Delete ACL.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.ACL().Delete()
	},
}

func init() {
	iamCmd.AddCommand(iamACLsCmd)
	iamACLsCmd.AddCommand(iamACLsListCmd)
	iamACLsCmd.AddCommand(iamACLsShowCmd)
	iamACLsCmd.AddCommand(iamACLsSetCmd)
	iamACLsCmd.AddCommand(iamACLsDeleteCmd)

	iamACLsCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
}
