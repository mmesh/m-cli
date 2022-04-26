package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// networkRoutesCmd represents the networkRoutes command
var networkRoutesCmd = &cobra.Command{
	Use:   "routes",
	Short: "Show network routes",
	Long:  appHeader(`Show network routes.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.RoutingTable().List()
	},
}

func init() {
	networkRoutesCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	networkRoutesCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	networkRoutesCmd.PersistentFlags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
	networkRoutesCmd.PersistentFlags().StringVarP(&vars.VRFID, "vrf", "v", "", "subnet/vrf identifier")
}
