package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// networkNodeCmd represents the networkNode command
var networkNodeCmd = &cobra.Command{
	Use:   "node",
	Short: "Network node operations",
	Long:  `Network node operations for network administrators.`,
}

// networkNodeCreateGenericNodeCmd represents the node create verb
var networkAddNodeCmd = &cobra.Command{
	Use:   "add",
	Short: "Generate magic link to setup a new node",
	Long:  `Generate magic link to setup a new node.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().AddNode()
	},
}

// networkNodeListCmd represents the admin/networks/vrfs list verb
var networkNodeListCmd = &cobra.Command{
	Use:   "list",
	Short: "List nodes",
	Long:  `List nodes.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().List()
	},
}

// networkNodeShowCmd represents the network/node show verb
var networkNodeShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show node",
	Long:  `Show node details.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().Show()
	},
}

// networkNodeDeleteCmd represents the network/nodes delete verb
var networkNodeDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove node",
	Long:  `Remove node from database.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().Delete()
	},
}

// networkNodeMetricsCmd represents the network/node metrics verb
var networkNodeMetricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Show detailed metrics",
	Long:  `Show detailed metrics.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().Metrics()
	},
}

// networkNodeResetTrafficMatrixCmd represents the network/nodes delete verb
// var networkNodeResetTrafficMatrixCmd = &cobra.Command{
// 	Use:   "reset-traffic-matrix",
// 	Short: "Reset traffic matrix metrics",
// 	Long:  `Reset traffic matrix metrics.`,
// 	Args:  cobra.NoArgs,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		client.Node().ResetNetworkTraffic()
// 	},
// }

// networkNodeShowEndpointCmd represents the network/node show-endpoint verb
var networkNodeShowEndpointCmd = &cobra.Command{
	Use:   "show-endpoint",
	Short: "Show network endpoint details",
	Long:  `Show network endpoint details.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().ShowEndpoint()
	},
}

// networkNodeDeleteEndpointCmd represents the network/node delete-endpoint verb
var networkNodeDeleteEndpointCmd = &cobra.Command{
	Use:   "delete-endpoint",
	Short: "Delete network endpoint",
	Long:  `Remove network endpoint from database.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Node().DeleteEndpoint()
	},
}

func init() {
	// rootCmd.AddCommand(networkNodeCmd)
	networkNodeCmd.AddCommand(networkAddNodeCmd)
	networkNodeCmd.AddCommand(networkNodeListCmd)
	networkNodeCmd.AddCommand(networkNodeShowCmd)
	networkNodeCmd.AddCommand(networkNodeDeleteCmd)
	networkNodeCmd.AddCommand(networkNodeMetricsCmd)
	// networkNodeCmd.AddCommand(networkNodeResetTrafficMatrixCmd)
	networkNodeCmd.AddCommand(networkNodeShowEndpointCmd)
	networkNodeCmd.AddCommand(networkNodeDeleteEndpointCmd)

	networkNodeCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	networkNodeCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	networkNodeCmd.PersistentFlags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
	networkNodeCmd.PersistentFlags().StringVarP(&vars.VRFID, "subnet", "s", "", "subnet/vrf identifier")
	networkNodeCmd.PersistentFlags().StringVarP(&vars.NodeID, "node", "x", "", "node identifier")
}
