package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/vars"
)

// supportCmd represents the support command
var supportCmd = &cobra.Command{
	Use:   "support",
	Short: "Get support and share your feedback",
	Long:  appHeader(`Get support and share your feedback.`),
}

// supportLiveChatCmd represents the support/chat verb
var supportLiveChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Open a live chat session",
	Long:  appHeader(`Open a live chat session.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.ITSM().Chat()
	},
}

// supportIssueCmd represents the support/issue command
var supportIssueCmd = &cobra.Command{
	Use:   "ticket",
	Short: "Manage your support tickets",
	Long:  appHeader(`Manage your support tickets.`),
}

// supportIssueNewCmd represents the support/issue create verb
var supportIssueNewCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new ticket",
	Long:  appHeader(`Create a new ticket.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.ITSM().Issue().New()
	},
}

// supportIssueListCmd represents the support/issue list verb
var supportIssueListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your tickets",
	Long:  appHeader(`List all your tickets.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.ITSM().Issue().List()
	},
}

// supportIssueShowCmd represents the support/issue get verb
var supportIssueShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Browse ticket",
	Long:  appHeader(`Show ticket details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		i := client.ITSM().Issue().Show()

		if input.GetConfirm("Add comment to this ticket?", false) {
			client.ITSM().Issue().NewComment(i)
		} else {
			fmt.Println()
		}
	},
}

// var supportIssueAddCommentCmd = &cobra.Command{
// 	Use:   "comment",
// 	Short: "Add a new comment to existing ticket",
// 	Long:  appHeader(`Add a new comment to existing ticket.`),
// 	Args:  cobra.NoArgs,
//	PreRun: func(cmd *cobra.Command, args []string) {
//		preflight()
//	},
// 	Run: func(cmd *cobra.Command, args []string) {
// 		client.ITSM().Issue().NewComment(nil)
// 	},
// }

// supportIssueCloseCmd represents the support/issue close verb
var supportIssueCloseCmd = &cobra.Command{
	Use:   "close",
	Short: "Mark ticket as 'closed'",
	Long:  appHeader(`Mark ticket as 'closed'.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.ITSM().Issue().Close()
	},
}

// supportIssueDeleteCmd represents the support/issue delete verb
var supportIssueDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete ticket",
	Long:  appHeader(`Remove ticket from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.ITSM().Issue().Delete()
	},
}

func init() {
	supportCmd.AddCommand(supportLiveChatCmd)
	supportCmd.AddCommand(supportIssueCmd)
	supportIssueCmd.AddCommand(supportIssueNewCmd)
	supportIssueCmd.AddCommand(supportIssueListCmd)
	supportIssueCmd.AddCommand(supportIssueShowCmd)
	// supportIssueCmd.AddCommand(supportIssueAddCommentCmd)
	supportIssueCmd.AddCommand(supportIssueCloseCmd)
	supportIssueCmd.AddCommand(supportIssueDeleteCmd)

	supportCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
}
