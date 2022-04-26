package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-api-go/grpc/resources/services/catalog/cloud/compute"
	"mmesh.dev/m-cli/internal/app/cli/auth/login"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// cloudAppCmd represents the cloud/app command
var cloudAppCmd = &cobra.Command{
	Use:   "app",
	Short: "Cloud apps management",
	Long:  appHeader(`Cloud apps management.`),
}

// cloudAppAddCmd represents the cloud/app set verb
var cloudAppAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a cloud app to your mmesh",
	Long:  appHeader(`Add a cloud app to your mmesh interactively.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().Create(login.NewRequest(), compute.ImageClass_APP)
	},
}

// cloudAppListCmd represents the cloud/app list verb
var cloudAppListCmd = &cobra.Command{
	Use:   "list",
	Short: "List cloud apps",
	Long:  appHeader(`List all cloud apps.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().List(compute.ImageClass_APP)
	},
}

// cloudAppShowCmd represents the cloud/app get verb
var cloudAppShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show cloud app",
	Long:  appHeader(`Show cloud app details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().Show(compute.ImageClass_APP)
	},
}

// cloudAppDeleteCmd represents the cloud/app delete verb
var cloudAppDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove a cloud app from your mmesh",
	Long:  appHeader(`Remove a cloud app from your mmesh.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().Delete(compute.ImageClass_APP)
	},
}

// cloudAppPowerCycleCmd represents the cloud/app power-cycle verb
var cloudAppPowerCycleCmd = &cobra.Command{
	Use:   "power-cycle",
	Short: "Power cycle a cloud app",
	Long:  appHeader(`Power cycle a cloud app.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().PowerCycle(compute.ImageClass_APP)
	},
}

// cloudAppPowerOnCmd represents the cloud/app power-on verb
var cloudAppPowerOnCmd = &cobra.Command{
	Use:   "power-on",
	Short: "Power-On a cloud app",
	Long:  appHeader(`Power-On a cloud app.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().PowerOn(compute.ImageClass_APP)
	},
}

// cloudAppPowerOffCmd represents the cloud/app power-off verb
var cloudAppPowerOffCmd = &cobra.Command{
	Use:   "power-off",
	Short: "Power-Off a cloud app",
	Long:  appHeader(`Power-Off a cloud app.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().PowerOff(compute.ImageClass_APP)
	},
}

// cloudAppRebootCmd represents the cloud/app reboot verb
var cloudAppRebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "Reboot a cloud app",
	Long:  appHeader(`Reboot a cloud app.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().Reboot(compute.ImageClass_APP)
	},
}

// cloudAppShutdownCmd represents the cloud/app reboot verb
var cloudAppShutdownCmd = &cobra.Command{
	Use:   "shutdown",
	Short: "Shutdown a cloud app",
	Long:  appHeader(`Shutdown a cloud app.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().Cloud().Instance().Shutdown(compute.ImageClass_APP)
	},
}

func init() {
	cloudCmd.AddCommand(cloudAppCmd)
	cloudAppCmd.AddCommand(cloudAppAddCmd)
	cloudAppCmd.AddCommand(cloudAppListCmd)
	cloudAppCmd.AddCommand(cloudAppShowCmd)
	cloudAppCmd.AddCommand(cloudAppDeleteCmd)
	cloudAppCmd.AddCommand(cloudAppPowerCycleCmd)
	cloudAppCmd.AddCommand(cloudAppPowerOnCmd)
	cloudAppCmd.AddCommand(cloudAppPowerOffCmd)
	cloudAppCmd.AddCommand(cloudAppRebootCmd)
	cloudAppCmd.AddCommand(cloudAppShutdownCmd)

	cloudAppCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	cloudAppAddCmd.Flags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	cloudAppAddCmd.Flags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
	cloudAppAddCmd.Flags().StringVarP(&vars.VRFID, "subnet", "s", "", "subnet/vrf identifier")
	cloudAppAddCmd.Flags().StringVarP(&vars.NodeID, "node", "x", "", "node identifier")
}
