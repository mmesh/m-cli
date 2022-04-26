package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// opsOperationsCmd represents the opsOperations command
var opsOperationsCmd = &cobra.Command{
	Use:   "log",
	Short: "Workflow log administration",
	Long:  appHeader(`Workflow log administration.`),
}

// opsOperationsListCmd represents the ops/operations list verb
var opsOperationsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List logs",
	Long:  appHeader(`List all logs.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Operation().List()
	},
}

// opsOperationsShowCmd represents the ops/operations get verb
var opsOperationsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show log",
	Long:  appHeader(`Show log details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Operation().Show()
	},
}

// opsOperationsDeleteCmd represents the ops/operations delete verb
var opsOperationsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete log",
	Long:  appHeader(`Remove log from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Operation().Delete()
	},
}

func init() {
	opsCmd.AddCommand(opsOperationsCmd)
	opsOperationsCmd.AddCommand(opsOperationsListCmd)
	opsOperationsCmd.AddCommand(opsOperationsShowCmd)
	opsOperationsCmd.AddCommand(opsOperationsDeleteCmd)

	opsOperationsCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	opsOperationsCmd.PersistentFlags().StringVarP(&vars.ProjectID, "project", "p", "", "project identifier")
	opsOperationsCmd.PersistentFlags().StringVarP(&vars.WorkflowID, "workflow", "w", "", "workflow identifier")
	opsOperationsCmd.PersistentFlags().StringVarP(&vars.OperationID, "operation", "o", "", "operation identifier")
}
