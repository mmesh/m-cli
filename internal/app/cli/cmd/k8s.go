package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Connect kubernetes clusters to your mmesh",
	Long:  appHeader(`Connect kubernetes clusters to your mmesh.`),
}

var k8sServicesCmd = &cobra.Command{
	Use:   "svc",
	Short: "Connect kubernetes services to your mmesh",
	Long:  appHeader(`Connect kubernetes services to your mmesh.`),
}

var k8sWorkloadsCmd = &cobra.Command{
	Use:   "workload",
	Short: "Connect kubernetes workloads to your mmesh",
	Long:  appHeader(`Connect kubernetes workloads to your mmesh.`),
}

// k8sCreateKubernetesGatewayCmd represents the node create verb
// var k8sCreateKubernetesGatewayCmd = &cobra.Command{
// 	Use:   "add-gw",
// 	Short: "Add mmesh gateway to your kubernetes cluster",
// 	Long:  appHeader(`Add mmesh gateway to your kubernetes cluster.`),
// 	Args:  cobra.NoArgs,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		client.Kubernetes().CreateKubernetesGateway()
// 	},
// }

var k8sDeleteGatewayCmd = &cobra.Command{
	Use:   "delete-gw",
	Short: "Remove mmesh gateway from your kubernetes cluster",
	Long:  appHeader(`Remove mmesh gateway from your kubernetes cluster.`),
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Kubernetes().DeleteGateway()
	},
}

var k8sServicesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List kubernetes services connected via mmesh ingress gateway",
	Long:  appHeader(`List kubernetes services connected via mmesh ingress gateway.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Kubernetes().Services()
	},
}

var k8sServicesConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect kubernetes services via mmesh ingress gateway",
	Long:  appHeader(`Connect kubernetes services via mmesh ingress gateway.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Kubernetes().ConnectService()
	},
}

var k8sServicesDisconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Disconnect kubernetes services from mmesh ingress gateway",
	Long:  appHeader(`Disconnect kubernetes services from mmesh ingress gateway.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Kubernetes().DisconnectService()
	},
}

var k8sWorkloadsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List kubernetes workloads connected via mmesh sidecar",
	Long:  appHeader(`List kubernetes workloads connected via mmesh sidecar.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Kubernetes().Pods()
	},
}

var k8sWorkloadsConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Add mmesh sidecar to your kubernetes workloads",
	Long:  appHeader(`Add mmesh sidecar to your kubernetes workloads.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Kubernetes().ConnectPod()
	},
}

var k8sWorkloadsDisconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Remove mmesh sidecar from your kubernetes workloads",
	Long:  appHeader(`Remove mmesh sidecar from your kubernetes workloads.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Kubernetes().DisconnectPod()
	},
}

func init() {
	k8sCmd.AddCommand(k8sServicesCmd)
	k8sCmd.AddCommand(k8sWorkloadsCmd)
	k8sCmd.AddCommand(k8sDeleteGatewayCmd)
	k8sServicesCmd.AddCommand(k8sServicesListCmd)
	k8sServicesCmd.AddCommand(k8sServicesConnectCmd)
	k8sServicesCmd.AddCommand(k8sServicesDisconnectCmd)
	k8sWorkloadsCmd.AddCommand(k8sWorkloadsListCmd)
	k8sWorkloadsCmd.AddCommand(k8sWorkloadsConnectCmd)
	k8sWorkloadsCmd.AddCommand(k8sWorkloadsDisconnectCmd)

	k8sCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	k8sCmd.PersistentFlags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
	k8sCmd.PersistentFlags().StringVarP(&vars.SubnetID, "subnet", "s", "", "subnet identifier")
	// k8sCmd.PersistentFlags().StringVarP(&vars.NodeID, "node", "x", "", "node identifier")
}
