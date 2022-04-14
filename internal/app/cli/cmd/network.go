package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// networkCmd represents the network command
var networkCmd = &cobra.Command{
	Use:   "network",
	Short: "Networks administration",
	Long:  `Networks administration.`,
}

// networkListCmd represents the admin/networks list verb
var networkListCmd = &cobra.Command{
	Use:   "list",
	Short: "List networks",
	Long:  `List all networks.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Network().List()
	},
}

// networkShowCmd represents the admin/networks get verb
var networkShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show network",
	Long:  `Show network details.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Network().Show()
	},
}

// networkNewCmd represents the network create verb
var networkNewCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a network",
	Long:  `Create a network interactively.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Network().New()
	},
}

// networkUpdateCmd represents the network update verb
var networkUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a network",
	Long:  `Update a network interactively.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Network().Update()
	},
}

// networkDeleteCmd represents the admin/tenants delete verb
var networkDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove network",
	Long:  `Remove network from database.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Network().Delete()
	},
}

func init() {
	// topologyCmd.AddCommand(networkCmd)
	networkCmd.AddCommand(networkListCmd)
	networkCmd.AddCommand(networkShowCmd)
	networkCmd.AddCommand(networkNewCmd)
	networkCmd.AddCommand(networkUpdateCmd)
	networkCmd.AddCommand(networkDeleteCmd)

	networkCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	networkCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	networkCmd.PersistentFlags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
}
