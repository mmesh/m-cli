package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Account administration",
	Long:  appHeader(`Account administration for platform admins.`),
}

var accountShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show account details",
	Long:  appHeader(`Show all account details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().Show()
	},
}

var accountCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new account",
	Long:  appHeader(`Create new account.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		header()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().New()
	},
}

var accountSettingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Edit account settings and integrations",
	Long:  appHeader(`Edit account settings and integrations.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().Settings()
	},
}

var accountSubscriptionCmd = &cobra.Command{
	Use:   "subscription",
	Short: "Manage service subscription",
	Long:  appHeader(`Manage service subscription.`),
}

var accountSubscriptionShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show and update service subscription",
	Long:  appHeader(`Show and update service subscription.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().Subscription(nil, true)
	},
}

var accountSubscriptionApplyPromotionCmd = &cobra.Command{
	Use:   "promo-code",
	Short: "Apply promotion code",
	Long:  appHeader(`Apply promotion code.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().ApplyPromotion()
	},
}

var accountSubscriptionCancelCmd = &cobra.Command{
	Use:   "delete",
	Short: "Cancel subscription and delete account",
	Long:  appHeader(`Cancel subscription and delete account.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().Cancel()
	},
}

var billingCmd = &cobra.Command{
	Use:   "billing",
	Short: "Manage invoices and billable items",
	Long:  appHeader(`Manage invoices and billable items.`),
}

var accountBillingPortalCmd = &cobra.Command{
	Use:   "portal",
	Short: "Open the Billing Portal",
	Long:  appHeader(`Open the Billing Portal.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().BillingPortal(nil)
	},
}

func init() {
	accountCmd.AddCommand(accountShowCmd)
	accountCmd.AddCommand(accountCreateCmd)
	accountCmd.AddCommand(accountSettingsCmd)
	accountCmd.AddCommand(accountSubscriptionCmd)
	accountSubscriptionCmd.AddCommand(accountSubscriptionShowCmd)
	accountSubscriptionCmd.AddCommand(accountSubscriptionApplyPromotionCmd)
	accountSubscriptionCmd.AddCommand(accountSubscriptionCancelCmd)
	accountCmd.AddCommand(billingCmd)
	billingCmd.AddCommand(accountBillingPortalCmd)

	accountCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
}
