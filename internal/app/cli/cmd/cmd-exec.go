package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// execCmd represents the mmp cmdExec
var execCmd = &cobra.Command{
	Use:   "exec [flags] [-- <command>]",
	Short: "Execute command on target node",
	Long:  appHeader(`Execute command on target node`),
	Args:  cobra.MinimumNArgs(0),
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Command().Exec(args)
	},
}

func init() {
	execCmd.Flags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier (optional)")
	execCmd.Flags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier (optional)")
	execCmd.Flags().StringVarP(&vars.NetID, "network", "n", "", "network identifier (optional)")
	execCmd.Flags().StringVarP(&vars.VRFID, "subnet", "s", "", "subnet/vrf identifier (optional)")
	execCmd.Flags().StringVarP(&vars.NodeID, "node", "x", "", "node identifier (optional)")
}
