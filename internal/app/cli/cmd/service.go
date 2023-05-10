package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Marketplace of professional services for your mmesh",
	Long:  appHeader(`Marketplace of professional services for your mmesh.`),
}

// serviceCatalogListCmd represents the service/crm list verb
var serviceCatalogListCmd = &cobra.Command{
	Use:   "list",
	Short: "List product/services in mmesh marketplace",
	Long:  appHeader(`List product/services in mmesh marketplace.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Product().List(true)
	},
}

// serviceCatalogShowCmd represents the service/crm show verb
var serviceCatalogShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show product/service in mmesh marketplace",
	Long:  appHeader(`Show product/service in mmesh marketplace.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Product().Show(true)
	},
}

func init() {
	serviceCmd.AddCommand(serviceCatalogListCmd)
	serviceCmd.AddCommand(serviceCatalogShowCmd)
}
