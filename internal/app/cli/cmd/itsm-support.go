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
	Long:  `Get support and share your feedback.`,
}

// supportLiveChatCmd represents the support/chat verb
var supportLiveChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Open a live chat session",
	Long:  `Open a live chat session.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.ITSM().Chat()
	},
}

// supportAssistanceInfoCmd represents the support/info submit verb
// var supportAssistanceInfoCmd = &cobra.Command{
// 	Use:   "info",
// 	Short: "Ask for general information",
// 	Long:  `Ask for general information.`,
// 	Args:  cobra.NoArgs,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		client.ITSM().AssistanceInfo()
// 	},
// }

// supportAssistanceRequestCmd represents the support/info submit verb
// var supportAssistanceRequestCmd = &cobra.Command{
// 	Use:   "request",
// 	Short: "Request technical support about the platform or the service",
// 	Long:  `Request technical support about the platform or the service.`,
// 	Args:  cobra.NoArgs,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		client.ITSM().AssistanceSupport()
// 	},
// }

// supportFeedbackCmd represents the support/info submit verb
// var supportFeedbackCmd = &cobra.Command{
// 	Use:   "feedback",
// 	Short: "Share your comments and suggestions",
// 	Long:  `Share your comments and suggestions.`,
// 	Args:  cobra.NoArgs,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		client.ITSM().Feedback()
// 	},
// }

// supportIncidentCmd represents the support/info submit verb
// var supportIncidentCmd = &cobra.Command{
// 	Use:   "incident",
// 	Short: "Report an incident",
// 	Long:  `Report an incident.`,
// 	Args:  cobra.NoArgs,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		client.ITSM().Incident()
// 	},
// }

// supportIssueCmd represents the support/issue command
var supportIssueCmd = &cobra.Command{
	Use:   "ticket",
	Short: "Manage your support tickets",
	Long:  `Manage your support tickets.`,
}

// supportIssueNewCmd represents the support/issue create verb
var supportIssueNewCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new ticket",
	Long:  `Create a new ticket.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.ITSM().Issue().New()
	},
}

// supportIssueListCmd represents the support/issue list verb
var supportIssueListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your tickets",
	Long:  `List all your tickets.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.ITSM().Issue().List()
	},
}

// supportIssueShowCmd represents the support/issue get verb
var supportIssueShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Browse ticket",
	Long:  `Show ticket details.`,
	Args:  cobra.NoArgs,
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
// 	Long:  `Add a new comment to existing ticket.`,
// 	Args:  cobra.NoArgs,
// 	Run: func(cmd *cobra.Command, args []string) {
// 		client.ITSM().Issue().NewComment(nil)
// 	},
// }

// supportIssueCloseCmd represents the support/issue close verb
var supportIssueCloseCmd = &cobra.Command{
	Use:   "close",
	Short: "Mark ticket as 'closed'",
	Long:  `Mark ticket as 'closed'.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.ITSM().Issue().Close()
	},
}

// supportIssueDeleteCmd represents the support/issue delete verb
var supportIssueDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete ticket",
	Long:  `Remove ticket from database.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.ITSM().Issue().Delete()
	},
}

func init() {
	supportCmd.AddCommand(supportLiveChatCmd)
	// supportCmd.AddCommand(supportAssistanceInfoCmd)
	// supportCmd.AddCommand(supportAssistanceRequestCmd)
	// supportCmd.AddCommand(supportFeedbackCmd)
	// supportCmd.AddCommand(supportIncidentCmd)
	supportCmd.AddCommand(supportIssueCmd)
	supportIssueCmd.AddCommand(supportIssueNewCmd)
	supportIssueCmd.AddCommand(supportIssueListCmd)
	supportIssueCmd.AddCommand(supportIssueShowCmd)
	// supportIssueCmd.AddCommand(supportIssueAddCommentCmd)
	supportIssueCmd.AddCommand(supportIssueCloseCmd)
	supportIssueCmd.AddCommand(supportIssueDeleteCmd)

	supportCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
}
