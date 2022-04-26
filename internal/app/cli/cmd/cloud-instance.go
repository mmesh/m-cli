package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-api-go/grpc/resources/services/catalog/cloud/compute"
	"mmesh.dev/m-cli/internal/app/cli/auth/login"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// cloudInstanceCmd represents the cloud/instance command
var cloudInstanceCmd = &cobra.Command{
	Use:   "instance",
	Short: "Cloud instances management",
	Long:  appHeader(`Cloud instances management.`),
}

// cloudInstanceAddCmd represents the cloud/instance set verb
var cloudInstanceAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a cloud instance to your mmesh",
	Long:  appHeader(`Add a cloud instance to your mmesh interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().Create(login.NewRequest(), compute.ImageClass_OS)
	},
}

// cloudInstanceListCmd represents the cloud/instance list verb
var cloudInstanceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List cloud instances",
	Long:  appHeader(`List all cloud instances.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().List(compute.ImageClass_OS)
	},
}

// cloudInstanceShowCmd represents the cloud/instance get verb
var cloudInstanceShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show cloud instance",
	Long:  appHeader(`Show cloud instance details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().Show(compute.ImageClass_OS)
	},
}

// cloudInstanceDeleteCmd represents the cloud/instance delete verb
var cloudInstanceDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Destroy and remove a cloud instance from your mmesh",
	Long:  appHeader(`Destroy and remove a cloud instance from your mmesh.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().Delete(compute.ImageClass_OS)
	},
}

// cloudInstancePowerCycleCmd represents the cloud/instance power-cycle verb
var cloudInstancePowerCycleCmd = &cobra.Command{
	Use:   "power-cycle",
	Short: "Power cycle a cloud instance",
	Long:  appHeader(`Power cycle a cloud instance.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().PowerCycle(compute.ImageClass_OS)
	},
}

// cloudInstancePowerOnCmd represents the cloud/instance power-on verb
var cloudInstancePowerOnCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Power-On a cloud instance",
	Long:  appHeader(`Power-On a cloud instance.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().PowerOn(compute.ImageClass_OS)
	},
}

// cloudInstancePowerOffCmd represents the cloud/instance power-off verb
var cloudInstancePowerOffCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Power-Off a cloud instance",
	Long:  appHeader(`Power-Off a cloud instance.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().PowerOff(compute.ImageClass_OS)
	},
}

// cloudInstanceRebootCmd represents the cloud/instance reboot verb
var cloudInstanceRebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot a cloud instance",
	Long:  appHeader(`Reboot a cloud instance.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().Reboot(compute.ImageClass_OS)
	},
}

// cloudInstanceShutdownCmd represents the cloud/instance reboot verb
var cloudInstanceShutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Shutdown a cloud instance",
	Long:  appHeader(`Shutdown a cloud instance.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().Shutdown(compute.ImageClass_OS)
	},
}

func init() {
	cloudCmd.AddCommand(cloudInstanceCmd)
	cloudInstanceCmd.AddCommand(cloudInstanceAddCmd)
	cloudInstanceCmd.AddCommand(cloudInstanceListCmd)
	cloudInstanceCmd.AddCommand(cloudInstanceShowCmd)
	cloudInstanceCmd.AddCommand(cloudInstanceDeleteCmd)
	cloudInstanceCmd.AddCommand(cloudInstancePowerCycleCmd)
	cloudInstanceCmd.AddCommand(cloudInstancePowerOnCmd)
	cloudInstanceCmd.AddCommand(cloudInstancePowerOffCmd)
	cloudInstanceCmd.AddCommand(cloudInstanceRebootCmd)
	cloudInstanceCmd.AddCommand(cloudInstanceShutdownCmd)

	cloudInstanceCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	cloudInstanceAddCmd.Flags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	cloudInstanceAddCmd.Flags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
	cloudInstanceAddCmd.Flags().StringVarP(&vars.VRFID, "subnet", "s", "", "subnet/vrf identifier")
	cloudInstanceAddCmd.Flags().StringVarP(&vars.NodeID, "node", "x", "", "node identifier")
}
