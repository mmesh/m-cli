package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// alertsCmd represents the alert command
var alertsCmd = &cobra.Command{
	Use:   "alert",
	Short: "Alert system",
	Long:  appHeader(`Alert system.`),
}

// alertAddCommentCmd represents the alert newComment verb
var alertAddCommentCmd = &cobra.Command{
	Use:   "comment",
	Short: "Add a new comment to alert",
	Long:  appHeader(`Add a new comment to alert.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Alert().NewComment()
	},
}

// alertListCmd represents the alert list verb
var alertListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your alerts",
	Long:  appHeader(`List all your alerts.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Alert().List()
	},
}

// alertShowCmd represents the alert get verb
var alertShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show alert",
	Long:  appHeader(`Show alert details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Alert().Show()
	},
}

// alertDeleteCmd represents the alert delete verb
var alertDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete alert",
	Long:  appHeader(`Remove alert from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Alert().Delete()
	},
}

func init() {
	alertsCmd.AddCommand(alertAddCommentCmd)
	alertsCmd.AddCommand(alertListCmd)
	alertsCmd.AddCommand(alertShowCmd)
	alertsCmd.AddCommand(alertDeleteCmd)

	alertsCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
}
