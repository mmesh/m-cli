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
	Long:  `Account administration for platform admins.`,
}

// accountCreateCmd represents the account create verb
var accountCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create new account",
	Long:  `Create new account.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().New()
	},
}

// accountSettingsCmd represents the account settings verb
var accountSettingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Edit account settings and integrations",
	Long:  `Edit account settings and integrations.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().Settings()
	},
}

// accountShowCmd represents the account get verb
var accountShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show all account details",
	Long:  `Show all account details.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().Show()
	},
}

// accountStatsCmd represents the account stats verb
var accountStatsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show account stats and usage",
	Long:  `Show account stats and usage.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().Stats()
	},
}

// accountSubscriptionCmd represents the billing command
var accountSubscriptionCmd = &cobra.Command{
	Use:   "subscription",
	Short: "Manage service subscription",
	Long:  `Manage service subscription.`,
}

// accountSubscriptionShowCmd represents the account subscription show verb
var accountSubscriptionShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show and update service subscription",
	Long:  `Show and update service subscription.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().Subscription()
	},
}

// accountSubscriptionApplyPromotionCmd represents the account subscription promo-code verb
var accountSubscriptionApplyPromotionCmd = &cobra.Command{
	Use:   "promo-code",
	Short: "Apply promotion code",
	Long:  `Apply promotion code.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().ApplyPromotion()
	},
}

// accountSubscriptionCancelCmd represents the account subscription delete verb
var accountSubscriptionCancelCmd = &cobra.Command{
	Use:   "delete",
	Short: "Cancel subscription and delete account",
	Long:  `Cancel subscription and delete account.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Account().Cancel()
	},
}

// billingCmd represents the billing command
var billingCmd = &cobra.Command{
	Use:   "billing",
	Short: "Manage invoices and billable items",
	Long:  `Manage invoices and billable items.`,
}

func init() {
	accountCmd.AddCommand(accountCreateCmd)
	accountCmd.AddCommand(accountSettingsCmd)

	accountCmd.AddCommand(accountShowCmd)
	accountCmd.AddCommand(accountStatsCmd)
	accountCmd.AddCommand(accountSubscriptionCmd)
	accountSubscriptionCmd.AddCommand(accountSubscriptionShowCmd)
	accountSubscriptionCmd.AddCommand(accountSubscriptionApplyPromotionCmd)
	accountSubscriptionCmd.AddCommand(accountSubscriptionCancelCmd)

	accountCmd.AddCommand(billingCmd)

	accountCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
}
