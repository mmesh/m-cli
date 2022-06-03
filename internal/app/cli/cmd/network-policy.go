package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

var networkPolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Security policies administration",
	Long:  appHeader(`Security policies administration.`),
}

var networkPolicyShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show security policy",
	Long:  appHeader(`Show security policy details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().Show()
	},
}

var networkPolicyImportCmd = &cobra.Command{
	Use:   "import -f <yamlFile>",
	Short: "Import security policy from YAML file",
	Long:  appHeader(`Import security policy from YAML file.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().Import(vars.YAMLFile)
	},
}

var networkPolicyExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export security policy to standard output",
	Long:  appHeader(`Export security policy  to standard output.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().Export()
	},
}

var networkPolicySetDefaultCmd = &cobra.Command{
	Use:   "set-default",
	Short: "Set default policy",
	Long:  appHeader(`Set default policy.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().SetDefault()
	},
}

var networkPolicySetRuleCmd = &cobra.Command{
	Use:   "set-rule",
	Short: "Set security policy rule",
	Long:  appHeader(`Set security policy rule.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().SetRule()
	},
}

var networkPolicyUnsetRuleCmd = &cobra.Command{
	Use:   "unset-rule",
	Short: "Unset security policy rule",
	Long:  appHeader(`Unset security policy rule.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().UnsetRule()
	},
}

var networkPolicyDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete security policy (all rules)",
	Long:  appHeader(`Remove security policy (all rules) from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().Delete()
	},
}

func init() {
	networkPolicyCmd.AddCommand(networkPolicyShowCmd)
	networkPolicyCmd.AddCommand(networkPolicyImportCmd)
	networkPolicyCmd.AddCommand(networkPolicyExportCmd)
	networkPolicyCmd.AddCommand(networkPolicySetDefaultCmd)
	networkPolicyCmd.AddCommand(networkPolicySetRuleCmd)
	networkPolicyCmd.AddCommand(networkPolicyUnsetRuleCmd)
	networkPolicyCmd.AddCommand(networkPolicyDeleteCmd)

	networkPolicyCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")
	networkPolicyCmd.PersistentFlags().StringVarP(&vars.TenantID, "tenant", "t", "", "tenant identifier")
	networkPolicyCmd.PersistentFlags().StringVarP(&vars.NetID, "network", "n", "", "network identifier")
	networkPolicyCmd.PersistentFlags().StringVarP(&vars.VRFID, "vrf", "v", "", "subnet/vrf identifier")

	networkPolicyImportCmd.Flags().StringVarP(&vars.YAMLFile, "yamlFile", "f", "", "yaml workflow file")
	networkPolicyImportCmd.MarkFlagRequired("yamlFile")
}
