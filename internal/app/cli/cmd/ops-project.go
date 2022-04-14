package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// opsProjectsCmd represents the opsProjects command
var opsProjectsCmd = &cobra.Command{
	Use:   "project",
	Short: "Workflow projects administration",
	Long:  `Workflow projects administration.`,
}

// opsProjectsListCmd represents the ops/projects list verb
var opsProjectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List projects",
	Long:  `List all projects.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Project().List()
	},
}

// opsProjectsShowCmd represents the ops/projects get verb
var opsProjectsShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show project",
	Long:  `Show project details.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Project().Show()
	},
}

// opsProjectsSetCmd represents the ops/projects set verb
var opsProjectsSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Create or update a project",
	Long:  `Create or update a project interactively.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Project().Set()
	},
}

// opsProjectsDeleteCmd represents the ops/projects delete verb
var opsProjectsDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove project",
	Long:  `Remove project from database.`,
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		client.Project().Delete()
	},
}

func init() {
	opsCmd.AddCommand(opsProjectsCmd)
	opsProjectsCmd.AddCommand(opsProjectsListCmd)
	opsProjectsCmd.AddCommand(opsProjectsShowCmd)
	opsProjectsCmd.AddCommand(opsProjectsSetCmd)
	opsProjectsCmd.AddCommand(opsProjectsDeleteCmd)

	opsProjectsCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	opsProjectsCmd.PersistentFlags().StringVarP(&vars.ProjectID, "project", "p", "", "project identifier")
}
