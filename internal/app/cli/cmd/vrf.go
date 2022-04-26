package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// subnetCmd represents the subnet command
var subnetCmd = &cobra.Command{
	Use:   "subnet",
	Short: "Subnets administration",
	Long:  appHeader(`Subnets administration.`),
}

// subnetListCmd represents the subnet list verb
var subnetListCmd = &cobra.Command{
	Use:   "list",
	Short: "List subnets",
	Long:  appHeader(`List all subnets.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VRF().List()
	},
}

// subnetShowCmd represents the subnet show verb
var subnetShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show subnet",
	Long:  appHeader(`Show subnet details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VRF().Show()
	},
}

// subnetNewCmd represents the subnet create verb
var subnetNewCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new subnet",
	Long:  appHeader(`Create a new subnet interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VRF().New()
	},
}

// subnetUpdateCmd represents the subnet update verb
var subnetUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a subnet",
	Long:  appHeader(`Update a subnet interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VRF().Update()
	},
}

// subnetDeleteCmd represents the subnet delete verb
var subnetDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove subnet",
	Long:  appHeader(`Remove subnet from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VRF().Delete()
	},
}

// subnetDeleteIPAMEntryCmd represents the subnet delete-ipam-endpoint verb
var subnetDeleteIPAMEntryCmd = &cobra.Command{
	Use:   "delete-ipam-entry",
	Short: "Delete IPAM entry",
	Long:  appHeader(`Remove IPAM entry from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.VRF().DeleteIPAMEntry()
	},
}

func init() {
	subnetCmd.AddCommand(subnetListCmd)
	subnetCmd.AddCommand(subnetShowCmd)
	subnetCmd.AddCommand(subnetNewCmd)
	subnetCmd.AddCommand(subnetUpdateCmd)
	subnetCmd.AddCommand(subnetDeleteCmd)
	subnetCmd.AddCommand(subnetDeleteIPAMEntryCmd)

	subnetCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	subnetCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	subnetCmd.PersistentFlags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
	subnetCmd.PersistentFlags().StringVarP(&vars.VRFID, "subnet", "s", "", "subnet/vrf identifier")
}
