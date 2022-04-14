package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// billingInvoiceCmd represents the billingInvoice command
var billingInvoiceCmd = &cobra.Command{
	Use:   "invoice",
	Short: "Invoice administration",
	Long:  `Invoice administration.`,
}

// billingInvoiceListCmd represents the billing/invoices list verb
var billingInvoiceListCmd = &cobra.Command{
	Use:   "list",
	Short: "List invoices",
	Long:  `List all invoices.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Billing().Invoice().List()
	},
}

// billingInvoiceShowCmd represents the billing/invoices get verb
var billingInvoiceShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show invoice",
	Long:  `Show invoice details.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Billing().Invoice().Show()
	},
}

// billingInvoiceDeleteCmd represents the billing/invoices delete verb
var billingInvoiceDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete invoice",
	Long:  `Remove invoice from database.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Billing().Invoice().Delete()
	},
}

func init() {
	billingCmd.AddCommand(billingInvoiceCmd)
	billingInvoiceCmd.AddCommand(billingInvoiceListCmd)
	billingInvoiceCmd.AddCommand(billingInvoiceShowCmd)
	billingInvoiceCmd.AddCommand(billingInvoiceDeleteCmd)

	billingInvoiceCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	billingInvoiceCmd.PersistentFlags().StringVarP(&vars.InvoiceID, "invoice", "i", "", "invoice identifier")
}
