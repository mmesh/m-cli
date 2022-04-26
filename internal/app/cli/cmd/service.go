package cmd

import (
	"github.com/spf13/cobra"
	"mmesh.dev/m-cli/pkg/client"
	"mmesh.dev/m-cli/pkg/vars"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "Marketplace of professional services for your mmesh",
	Long:  appHeader(`Marketplace of professional services for your mmesh.`),
}

// serviceRequestCmd represents the service/request verb
var serviceRequestCmd = &cobra.Command{
	Use:   "request",
	Short: "Create RFS (Request for Service) for the mmesh marketplace",
	Long:  appHeader(`Create RFS (Request for Service) for the mmesh marketplace.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.ITSM().ServiceRequest()
	},
}

// serviceCatalogListCmd represents the service/crm list verb
var serviceCatalogListCmd = &cobra.Command{
	Use:   "list",
	Short: "List product/services in mmesh marketplace",
	Long:  appHeader(`List product/services in mmesh marketplace.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Product().List(true)
	},
}

// serviceCatalogShowCmd represents the service/crm show verb
var serviceCatalogShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show product/service in mmesh marketplace",
	Long:  appHeader(`Show product/service in mmesh marketplace.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Product().Show(true)
	},
}

// serviceCRMCmd represents the service/crm command
var serviceCRMCmd = &cobra.Command{
	Use:   "crm",
	Short: "Manage your service/product catalog and sales opportunities",
	Long:  appHeader(`Manage your service/product catalog and sales opportunities.`),
}

// serviceCRMOpportunityCmd represents the service/crm opportunity command
var serviceCRMOpportunityCmd = &cobra.Command{
	Use:   "opportunity",
	Short: "Manage your opportunities",
	Long:  appHeader(`Manage your opportunities.`),
}

// serviceCRMOpportunityListCmd represents the service/crm list verb
var serviceCRMOpportunityListCmd = &cobra.Command{
	Use:   "list",
	Short: "List your opportunities",
	Long:  appHeader(`List your opportunities.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().CRM().Opportunity().List()
	},
}

// serviceCRMOpportunityShowCmd represents the service/crm show verb
var serviceCRMOpportunityShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show opportunity",
	Long:  appHeader(`Show opportunity details.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().CRM().Opportunity().Show()
	},
}

// serviceCRMOpportunitySetMilestoneCmd represents the service/crm set-milestone verb
var serviceCRMOpportunitySetMilestoneCmd = &cobra.Command{
	Use:   "set-milestone",
	Short: "Set opportunity milestone",
	Long:  appHeader(`Set opportunity milestone.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().CRM().Opportunity().SetMilestone()
	},
}

// serviceCRMOpportunitySetOutcomeCmd represents the service/crm set-outcome verb
var serviceCRMOpportunitySetOutcomeCmd = &cobra.Command{
	Use:   "set-outcome",
	Short: "Set opportunity outcome",
	Long:  appHeader(`Set opportunity outcome.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().CRM().Opportunity().SetOutcome()
	},
}

// serviceCRMOpportunityDeleteCmd represents the service/crm delete verb
var serviceCRMOpportunityDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove opportunity from database",
	Long:  appHeader(`Remove opportunity from database.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Platform().CRM().Opportunity().Delete()
	},
}

// serviceCRMCatalogCmd represents the service/crm catalog command
var serviceCRMCatalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Manage your service/product catalog",
	Long:  appHeader(`Manage your service/product catalog.`),
}

// serviceCRMCatalogListCmd represents the service/crm list verb
var serviceCRMCatalogListCmd = &cobra.Command{
	Use:   "list",
	Short: "List product/services published in your catalog",
	Long:  appHeader(`List product/services published in your catalog.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Product().List(false)
	},
}

// serviceCRMCatalogShowCmd represents the service/crm show verb
var serviceCRMCatalogShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show product/service published in your catalog",
	Long:  appHeader(`Show product/service published in your catalog.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Product().Show(false)
	},
}

// serviceCRMCatalogSetCmd represents the service/crm setverb
var serviceCRMCatalogSetCmd = &cobra.Command{
	Use:   "publish -f <yamlFile>",
	Short: "Publish your service catalog in the mmesh marketplace",
	Long:  appHeader(`Publish your service catalog in the mmesh marketplace.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Product().Set(vars.YAMLFile)
	},
}

// serviceCRMCatalogDeleteCmd represents the service/crm delete verb
var serviceCRMCatalogDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove product/service from your catalog",
	Long:  appHeader(`Remove product/service from your catalog.`),
	Args:  cobra.NoArgs,
	PreRun: func(cmd *cobra.Command, args []string) {
		preflight()
	},
	Run: func(cmd *cobra.Command, args []string) {
		client.Services().Product().Delete()
	},
}

func init() {
	serviceCmd.AddCommand(serviceRequestCmd)
	serviceCmd.AddCommand(serviceCatalogListCmd)
	serviceCmd.AddCommand(serviceCatalogShowCmd)
	serviceCmd.AddCommand(serviceCRMCmd)
	serviceCRMCmd.AddCommand(serviceCRMOpportunityCmd)
	serviceCRMOpportunityCmd.AddCommand(serviceCRMOpportunityListCmd)
	serviceCRMOpportunityCmd.AddCommand(serviceCRMOpportunityShowCmd)
	serviceCRMOpportunityCmd.AddCommand(serviceCRMOpportunitySetMilestoneCmd)
	serviceCRMOpportunityCmd.AddCommand(serviceCRMOpportunitySetOutcomeCmd)
	serviceCRMOpportunityCmd.AddCommand(serviceCRMOpportunityDeleteCmd)
	serviceCRMCmd.AddCommand(serviceCRMCatalogCmd)
	serviceCRMCatalogCmd.AddCommand(serviceCRMCatalogListCmd)
	serviceCRMCatalogCmd.AddCommand(serviceCRMCatalogShowCmd)
	serviceCRMCatalogCmd.AddCommand(serviceCRMCatalogSetCmd)
	serviceCRMCatalogCmd.AddCommand(serviceCRMCatalogDeleteCmd)

	serviceCmd.PersistentFlags().StringVarP(&vars.AccountID, "account", "a", "", "account identifier")

	serviceCRMCatalogSetCmd.Flags().StringVarP(&vars.YAMLFile, "yamlFile", "f", "", "yaml service catalog file")
	serviceCRMCatalogSetCmd.MarkFlagRequired("yamlFile")
}
