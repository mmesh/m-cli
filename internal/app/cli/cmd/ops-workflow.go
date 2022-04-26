package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// opsWorkflowsCmd represents the opsWorkflows command
var opsWorkflowsCmd = &cobra.Command{
	Use:   "workflow",
	Short: "Automation workflows administration",
	Long:  appHeader(`Automation workflows administration.`),
}

// opsWorkflowsListCmd represents the ops/workflows list verb
var opsWorkflowsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List workflows",
	Long:  appHeader(`List all workflows.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().List()
	},
}

// opsWorkflowsShowCmd represents the ops/workflows get verb
var opsWorkflowsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show workflow",
	Long:  appHeader(`Show workflow details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().Show()
	},
}

// opsWorkflowsSetCmd represents the ops/workflows set verb
var opsWorkflowsSetCmd = &cobra.Command{
	Use:   "set -f <yamlFile>",
	Short: "Create or update workflow from YAML file",
	Long:  appHeader(`Create or update workflow from YAML file.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().Set(vars.YAMLFile)
	},
}

// opsWorkflowsDeleteCmd represents the ops/workflows delete verb
var opsWorkflowsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove workflow",
	Long:  appHeader(`Remove workflow from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().Delete()
	},
}

// opsWorkflowsEnableCmd represents the ops/workflows enable verb
var opsWorkflowsEnableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable workflow",
	Long:  appHeader(`Enable workflow.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().Enable()
	},
}

// opsWorkflowsDisableCmd represents the ops/workflows disable verb
var opsWorkflowsDisableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disable workflow",
	Long:  appHeader(`Disable workflow.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Workflow().Disable()
	},
}

func init() {
	opsCmd.AddCommand(opsWorkflowsCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsListCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsShowCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsSetCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsDeleteCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsEnableCmd)
	opsWorkflowsCmd.AddCommand(opsWorkflowsDisableCmd)

	opsWorkflowsCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	opsWorkflowsCmd.PersistentFlags().StringVarP(&vars.ProjectID, "project", "p", "", "project identifier")
	opsWorkflowsCmd.PersistentFlags().StringVarP(&vars.WorkflowID, "workflow", "w", "", "workflow identifier")

	opsWorkflowsSetCmd.Flags().StringVarP(&vars.YAMLFile, "yamlFile", "f", "", "yaml workflow file")
	opsWorkflowsSetCmd.MarkFlagRequired("yamlFile")
}
