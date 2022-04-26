package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// cloudKubernetesClusterCmd represents the cloud/kubernetes command
var cloudKubernetesClusterCmd = &cobra.Command{
	Use:   "kubernetes",
	Short: "Kubernetes clusters management",
	Long:  appHeader(`Kubernetes clusters management.`),
}

// cloudKubernetesClusterListCmd represents the cloud/kubernetes list verb
var cloudKubernetesClusterListCmd = &cobra.Command{
	Use:   "list",
	Short: "List kubernetes clusters",
	Long:  appHeader(`List all kubernetes clusters.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Kubernetes().Cluster().List()
	},
}

// cloudKubernetesClusterShowCmd represents the cloud/kubernetes get verb
var cloudKubernetesClusterShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show kubernetes cluster",
	Long:  appHeader(`Show kubernetes cluster details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Kubernetes().Cluster().Show()
	},
}

// cloudKubernetesClusterAddCmd represents the cloud/kubernetes set verb
var cloudKubernetesClusterAddCmd = &cobra.Command{
	Use:   "create",
	Short: "Create and add a kubernetes cluster to your mmesh",
	Long:  appHeader(`Create and add a kubernetes cluster to your mmesh interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Kubernetes().Cluster().Create()
	},
}

// cloudKubernetesClusterDeleteCmd represents the cloud/kubernetes delete verb
var cloudKubernetesClusterDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Destroy and remove a kubernetes cluster from your mmesh",
	Long:  appHeader(`Destroy and remove a kubernetes cluster from your mmesh.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Kubernetes().Cluster().Delete()
	},
}

// cloudKubernetesClusterGetKubeConfigCmd represents the cloud/kubernetes getKubeconfig verb
var cloudKubernetesClusterGetKubeConfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Get kubeconfig file in YAML format from cluster",
	Long:  appHeader(`Get kubeconfig file in YAML format from cluster.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Kubernetes().Cluster().GetKubeConfig()
	},
}

// cloudKubernetesCreateNodePoolCmd represents the cloud/kubernetes add-nodepool verb
var cloudKubernetesCreateNodePoolCmd = &cobra.Command{
	Use:   "add-nodepool",
	Short: "Add a new nodepool to an existing cluster",
	Long:  appHeader(`Add a new nodepool to an existing cluster.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Kubernetes().Cluster().CreateNodePool()
	},
}

// cloudKubernetesDeleteNodePoolCmd represents the cloud/kubernetes delete-nodepool verb
var cloudKubernetesDeleteNodePoolCmd = &cobra.Command{
	Use:   "delete-nodepool",
	Short: "Delete nodepool from kubernetes cluster",
	Long:  appHeader(`Delete nodepool from kubernetes cluster.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Kubernetes().Cluster().DeleteNodePool()
	},
}

// cloudKubernetesAddNodeCmd represents the cloud/kubernetes add-node verb
var cloudKubernetesAddNodeCmd = &cobra.Command{
	Use:   "add-node",
	Short: "Add a new node to an existing cluster",
	Long:  appHeader(`Add a new node to an existing cluster.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Kubernetes().Cluster().AddNode()
	},
}

// cloudKubernetesDeleteNodeCmd represents the cloud/kubernetes delete-nodepool verb
var cloudKubernetesDeleteNodeCmd = &cobra.Command{
	Use:   "delete-node",
	Short: "Delete node from kubernetes cluster",
	Long:  appHeader(`Delete node from kubernetes cluster.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Kubernetes().Cluster().DeleteNode()
	},
}

func init() {
	cloudCmd.AddCommand(cloudKubernetesClusterCmd)
	cloudKubernetesClusterCmd.AddCommand(cloudKubernetesClusterListCmd)
	cloudKubernetesClusterCmd.AddCommand(cloudKubernetesClusterShowCmd)
	cloudKubernetesClusterCmd.AddCommand(cloudKubernetesClusterAddCmd)
	cloudKubernetesClusterCmd.AddCommand(cloudKubernetesClusterDeleteCmd)
	cloudKubernetesClusterCmd.AddCommand(cloudKubernetesClusterGetKubeConfigCmd)
	cloudKubernetesClusterCmd.AddCommand(cloudKubernetesCreateNodePoolCmd)
	cloudKubernetesClusterCmd.AddCommand(cloudKubernetesDeleteNodePoolCmd)
	cloudKubernetesClusterCmd.AddCommand(cloudKubernetesAddNodeCmd)
	cloudKubernetesClusterCmd.AddCommand(cloudKubernetesDeleteNodeCmd)

	cloudKubernetesClusterCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	cloudKubernetesClusterAddCmd.Flags().StringVarP(&vars.CloudKubernetesClusterID, "cluster", "k", "", "kubernetes cluster identifier")
}
