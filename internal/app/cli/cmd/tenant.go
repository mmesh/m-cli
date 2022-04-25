package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// tenantCmd represents the tenant command
var tenantCmd = &cobra.Command{
	Use:   "tenant",
	Short: "Tenants administration",
	Long:  `Tenants administration.`,
}

// tenantListCmd represents the admin/tenants list verb
var tenantListCmd = &cobra.Command{
	Use:   "list",
	Short: "List tenants",
	Long:  `List all tenants.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(colors.Black(version.CLI_NAME + " " + version.GetVersion()))
		// output.AppHeader(vars.AccountID, false)
		client.Tenant().List()
	},
}

// tenantShowCmd represents the admin/tenants get verb
var tenantShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show tenant",
	Long:  `Show tenant details.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println(colors.Black(version.CLI_NAME + " " + version.GetVersion()))
		// output.AppHeader(vars.AccountID, false)
		client.Tenant().Show()
	},
}

// tenantSetCmd represents the admin/tenants set verb
var tenantSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Create or update tenant",
	Long:  `Create or update tenant interactively.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Tenant().Set()
	},
}

// tenantDeleteCmd represents the admin/tenants delete verb
var tenantDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove tenant",
	Long:  `Remove tenant from database.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Tenant().Delete()
	},
}

func init() {
	// topologyCmd.AddCommand(tenantCmd)
	tenantCmd.AddCommand(tenantListCmd)
	tenantCmd.AddCommand(tenantShowCmd)
	tenantCmd.AddCommand(tenantSetCmd)
	tenantCmd.AddCommand(tenantDeleteCmd)

	tenantCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	tenantCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
}
