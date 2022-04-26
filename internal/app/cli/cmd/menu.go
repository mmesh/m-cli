package cmd

import (
	"github.com/spf13/cobra"
)

// iamCmd represents the iam command
var iamCmd = &cobra.Command{
	Use:   "iam",
	Short: "Users administration and access management (RBAC)",
	Long:  appHeader(`Users administration and access management (RBAC).`),
}

// userCmd represents the iam command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "User settings",
	Long:  appHeader(`User settings.`),
}

// cloudCmd represents the k8s command
var cloudCmd = &cobra.Command{
	Use:   "cloud",
	Short: "Integrate and manage cloud resources in mmesh",
	Long:  appHeader(`Integrate and manage cloud resources in mmesh.`),
}

// opsCmd represents the ops command
var opsCmd = &cobra.Command{
	Use:   "ops",
	Short: "Automation and GitOps: projects / workflows / logs",
	Long:  appHeader(`Automation and workflows management commands.`),
}

// auditCmd represents the audit command
var auditCmd = &cobra.Command{
	Use:   "audit",
	Short: "Audit commands",
	Long:  appHeader(`Platform audit commands.`),
}

// statusCmd represents the status command
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Platform status report",
	Long:  appHeader(`Platform status report.`),
}

// runtimeCmd represents the runtime command
var runtimeCmd = &cobra.Command{
	Use:   "runtime",
	Short: "Runtime configuration administration",
	Long:  appHeader(`Runtime configuration administration.`),
}

func init() {
	rootCmd.AddCommand(accountCmd)
	rootCmd.AddCommand(iamCmd)
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(userCmd)
	rootCmd.AddCommand(alertsCmd)
	rootCmd.AddCommand(tenantCmd)
	rootCmd.AddCommand(networkCmd)
	rootCmd.AddCommand(subnetCmd)
	rootCmd.AddCommand(networkRoutesCmd)
	rootCmd.AddCommand(networkNodeCmd)
	rootCmd.AddCommand(networkPolicyCmd)
	rootCmd.AddCommand(k8sCmd)
	rootCmd.AddCommand(cloudCmd)
	rootCmd.AddCommand(serviceCmd)
	rootCmd.AddCommand(opsCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(transferCmd)
	rootCmd.AddCommand(portFwdCmd)
	// rootCmd.AddCommand(auditCmd)
	rootCmd.AddCommand(supportCmd)
	// rootCmd.AddCommand(statusCmd)
	// rootCmd.AddCommand(runtimeCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(setupCmd)
	rootCmd.AddCommand(completionCmd)
}
