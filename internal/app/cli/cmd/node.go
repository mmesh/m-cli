package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Network node operations",
	Long:  appHeader(`Network node operations for network administrators.`),
}

var nodeAddNodeCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new node",
	Long:  appHeader(`Register a new node.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().AddNode()
	},
}

/*
var nodeGetInstallationWebhookCmd = &cobra.Command{
	Use:   "get-magic-link",
	Short: "Generate magic link to setup a new node (linux only)",
	Long:  appHeader(`Generate magic link to setup a new node (linux only).`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().GetInstallationWebhook()
	},
}
*/

// nodeListCmd represents the admin/networks/vrfs list verb
var nodeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List nodes",
	Long:  appHeader(`List nodes.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().List()
	},
}

// nodeShowCmd represents the network/node show verb
var nodeShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show node",
	Long:  appHeader(`Show node details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().Show()
	},
}

// nodeDeleteCmd represents the network/nodes delete verb
var nodeDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove node",
	Long:  appHeader(`Remove node from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().Delete()
	},
}

// nodeMetricsCmd represents the network/node metrics verb
var nodeMetricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Show detailed metrics",
	Long:  appHeader(`Show detailed metrics.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().Metrics()
	},
}

// nodeResetTrafficMatrixCmd represents the network/nodes delete verb
// var nodeResetTrafficMatrixCmd = &cobra.Command{
// 	Use:   "reset-traffic-matrix",
// 	Short: "Reset traffic matrix metrics",
// 	Long:  appHeader(`Reset traffic matrix metrics.`),
// 	Args:  cobra.NoArgs,
//  PreRun: func(cmd *cobra.Command, args []string) {
// 	    preflight()
//  },
// 	Run: func(cmd *cobra.Command, args []string) {
// 		client.Node().ResetNetworkTraffic()
// 	},
// }

// nodeShowEndpointCmd represents the network/node show-endpoint verb
var nodeShowEndpointCmd = &cobra.Command{
	Use:   "show-endpoint",
	Short: "Show network endpoint details",
	Long:  appHeader(`Show network endpoint details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().ShowEndpoint()
	},
}

// nodeDeleteEndpointCmd represents the network/node delete-endpoint verb
var nodeDeleteEndpointCmd = &cobra.Command{
	Use:   "delete-endpoint",
	Short: "Delete network endpoint",
	Long:  appHeader(`Remove network endpoint from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().DeleteEndpoint()
	},
}

func init() {
	nodeCmd.AddCommand(nodeAddNodeCmd)
	// nodeCmd.AddCommand(nodeGetInstallationWebhookCmd)
	nodeCmd.AddCommand(nodeListCmd)
	nodeCmd.AddCommand(nodeShowCmd)
	nodeCmd.AddCommand(nodeDeleteCmd)
	nodeCmd.AddCommand(nodeMetricsCmd)
	// nodeCmd.AddCommand(nodeResetTrafficMatrixCmd)
	nodeCmd.AddCommand(nodeShowEndpointCmd)
	nodeCmd.AddCommand(nodeDeleteEndpointCmd)

	nodeCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	nodeCmd.PersistentFlags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
	nodeCmd.PersistentFlags().StringVarP(&vars.SubnetID, "subnet", "s", "", "subnet identifier")
	nodeCmd.PersistentFlags().StringVarP(&vars.NodeID, "node", "x", "", "node identifier")
}
