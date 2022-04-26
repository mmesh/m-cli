package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// networkPolicyCmd represents the networkPolicies command
var networkPolicyCmd = &cobra.Command{
	Use:   "policy",
	Short: "Network policies administration",
	Long:  appHeader(`Network policies operations for network administrators.`),
}

// networkPolicyShowCmd represents the network/policies get verb
var networkPolicyShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show network policy",
	Long:  appHeader(`Show network policy details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().Show()
	},
}

// networkPolicyImportCmd represents the network/endpoint import-policy verb
var networkPolicyImportCmd = &cobra.Command{
	Use:   "import -f <yamlFile>",
	Short: "Import network policy from YAML file",
	Long:  appHeader(`Import network policy from YAML file.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().Import(vars.YAMLFile)
	},
}

// networkPolicyExportCmd represents the network/endpoint import-policy verb
var networkPolicyExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export network policy to standard output",
	Long:  appHeader(`Export network policy  to standard output.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().Export()
	},
}

// networkPolicySetDefaultCmd represents the network/policy set-default verb
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

// networkPolicySetRuleCmd represents the network/endpoint set-rule verb
var networkPolicySetRuleCmd = &cobra.Command{
	Use:   "set-rule",
	Short: "Set network policy rule",
	Long:  appHeader(`Set network policy rule.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().SetRule()
	},
}

// networkPolicyUnsetRuleCmd represents the network/endpoint unset-policy verb
var networkPolicyUnsetRuleCmd = &cobra.Command{
	Use:   "unset-rule",
	Short: "Unset network policy rule",
	Long:  appHeader(`Unset network policy rule.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.NetworkPolicy().UnsetRule()
	},
}

// networkPolicyDeleteCmd represents the network/endpoint delete verb
var networkPolicyDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete network policy (all rules)",
	Long:  appHeader(`Remove network policy (all rules) from database.`),
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
