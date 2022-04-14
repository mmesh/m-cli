package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// iamSecurityGroupsCmd represents the iamSecurityGroups command
var iamSecurityGroupsCmd = &cobra.Command{
	Use:   "security-group",
	Short: "IAM security-groups adminstration",
	Long:  `IAM security-groups adminstration operations.`,
}

// iamSecurityGroupsListCmd represents the iam/security-groups list verb
var iamSecurityGroupsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List security-groups",
	Long:  `List all realm's security-groups.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.SecurityGroup().List()
	},
}

// iamSecurityGroupsShowCmd represents the iam/security-groups get verb
var iamSecurityGroupsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show security-group",
	Long:  `Show security-group details.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.SecurityGroup().Show()
	},
}

// iamSecurityGroupsSetCmd represents the iam/security-groups set verb
var iamSecurityGroupsSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Create or update security-group",
	Long:  `Create or update security-group.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.SecurityGroup().Set()
	},
}

// iamSecurityGroupsDeleteCmd represents the iam/security-groups delete verb
var iamSecurityGroupsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete security-group",
	Long:  `Delete security-group.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.SecurityGroup().Delete()
	},
}

func init() {
	iamCmd.AddCommand(iamSecurityGroupsCmd)
	iamSecurityGroupsCmd.AddCommand(iamSecurityGroupsListCmd)
	iamSecurityGroupsCmd.AddCommand(iamSecurityGroupsShowCmd)
	iamSecurityGroupsCmd.AddCommand(iamSecurityGroupsSetCmd)
	iamSecurityGroupsCmd.AddCommand(iamSecurityGroupsDeleteCmd)

	iamSecurityGroupsCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
}
