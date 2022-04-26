package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// billingItemCmd represents the billingItem command
var billingItemCmd = &cobra.Command{
	Use:   "item",
	Short: "Billed Item administration",
	Long:  appHeader(`Billed Item administration.`),
}

// billingItemListCmd represents the billing/invoices list verb
var billingItemListCmd = &cobra.Command{
	Use:   "list",
	Short: "List billed items",
	Long:  appHeader(`List all your billed items.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Billing().Item().List()
	},
}

// billingItemShowCmd represents the billing/invoices get verb
var billingItemShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show billed item",
	Long:  appHeader(`Show billed item details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Billing().Item().Show()
	},
}

func init() {
	billingCmd.AddCommand(billingItemCmd)
	billingItemCmd.AddCommand(billingItemListCmd)
	billingItemCmd.AddCommand(billingItemShowCmd)

	billingItemCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	billingItemCmd.PersistentFlags().StringVarP(&vars.BilledItemID, "item", "i", "", "billed item identifier")
}
