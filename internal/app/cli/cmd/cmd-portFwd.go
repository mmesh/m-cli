package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// portFwdCmd represents the mmp cmdPortFwd
var portFwdCmd = &cobra.Command{
	// Use:   "port-forward <nodeID:port> <nodeID:port>",
	Use:   "port-fwd",
	Short: "Forward local TCP port to target node",
	Long:  appHeader(`Forward local TCP port to target node.`),
	// Args:  cobra.ExactArgs(2),
	Args: cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Command().PortFwd()
	},
}

func init() {
	portFwdCmd.Flags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier (optional)")
	portFwdCmd.Flags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier (optional)")
	portFwdCmd.Flags().StringVarP(&vars.NetID, "network", "n", "", "network identifier (optional)")
	portFwdCmd.Flags().StringVarP(&vars.VRFID, "subnet", "s", "", "subnet/vrf identifier (optional)")
	portFwdCmd.Flags().StringVarP(&vars.NodeID, "node", "x", "", "node identifier (optional)")
}
