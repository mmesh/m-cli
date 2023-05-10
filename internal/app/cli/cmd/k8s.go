package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// k8sCmd represents the k8s command
var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Connect kubernetes clusters to your mmesh",
	Long:  appHeader(`Connect kubernetes clusters to your mmesh.`),
}

// k8sServicesCmd represents the node list-services verb
var k8sServicesCmd = &cobra.Command{
	Use:   "svc",
	Short: "Connect kubernetes services to your mmesh",
	Long:  appHeader(`Connect kubernetes services to your mmesh.`),
}

// k8sPodsCmd represents the node list-services verb
var k8sPodsCmd = &cobra.Command{
	Use:   "pod",
	Short: "Connect kubernetes pods to your mmesh",
	Long:  appHeader(`Connect kubernetes pods to your mmesh.`),
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

// k8sServicesListCmd represents the node list-services verb
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

// k8sServicesConnectCmd represents the node connect verb
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

// k8sServicesDisconnectCmd represents the node disconnect verb
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

// k8sPodsListCmd represents the node list-services verb
var k8sPodsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List kubernetes pods connected via mmesh sidecar",
	Long:  appHeader(`List kubernetes pods connected via mmesh sidecar.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Kubernetes().Pods()
	},
}

// k8sPodsConnectCmd represents the node add-sidecar verb
var k8sPodsConnectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Add mmesh sidecar to your kubernetes pods",
	Long:  appHeader(`Add mmesh sidecar to your kubernetes pods.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Kubernetes().ConnectPod()
	},
}

// k8sPodsDisconnectCmd represents the node remove-sidecar verb
var k8sPodsDisconnectCmd = &cobra.Command{
	Use:   "disconnect",
	Short: "Remove mmesh sidecar from your kubernetes pods",
	Long:  appHeader(`Remove mmesh sidecar from your kubernetes pods.`),
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
	k8sCmd.AddCommand(k8sPodsCmd)
	k8sServicesCmd.AddCommand(k8sServicesListCmd)
	k8sServicesCmd.AddCommand(k8sServicesConnectCmd)
	k8sServicesCmd.AddCommand(k8sServicesDisconnectCmd)
	k8sPodsCmd.AddCommand(k8sPodsListCmd)
	k8sPodsCmd.AddCommand(k8sPodsConnectCmd)
	k8sPodsCmd.AddCommand(k8sPodsDisconnectCmd)

	k8sCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	k8sCmd.PersistentFlags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
	k8sCmd.PersistentFlags().StringVarP(&vars.SubnetID, "subnet", "s", "", "subnet identifier")
	// k8sCmd.PersistentFlags().StringVarP(&vars.NodeID, "node", "x", "", "node identifier")
}
