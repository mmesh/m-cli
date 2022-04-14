package output

import (
	"fmt"
	"strings"

	"mmesh.dev/m-api-go/grpc/resources/services"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) ProductPrice(p *services.Product) string {
	var priceModel string
	var amount string

	switch p.PricingModel {
	case services.PricingModel_PRICING_FREE:
		return "Free"
	case services.PricingModel_PRICING_CUSTOM:
		return "Custom"
	case services.PricingModel_PRICING_FIXED:
		priceModel = "fixed"
		amount = fmt.Sprintf("$%v", p.PriceFixed)
	case services.PricingModel_PRICING_HOURLY:
		priceModel = "hourly"
		amount = fmt.Sprintf("$%v/hr", p.PriceHourly)
	case services.PricingModel_PRICING_WEEKLY:
		priceModel = "weekly"
		amount = fmt.Sprintf("$%v/wk", p.PriceWeekly)
	case services.PricingModel_PRICING_MONTHLY:
		priceModel = "monthly"
		amount = fmt.Sprintf("$%v/mo", p.PriceMonthly)
	case services.PricingModel_PRICING_YEARLY:
		priceModel = "yearly"
		amount = fmt.Sprintf("$%v/yr", p.PriceYearly)
	}

	if len(amount) == 0 {
		return "Custom"
	}

	return fmt.Sprintf("%s (%s)", amount, priceModel)
}

func (api *API) ProductSLA(sla services.ProductSLA) string {
	switch sla {
	case services.ProductSLA_BEST_EFFORT:
		return "Best Effort"
	case services.ProductSLA_BASIC:
		return "Basic"
	case services.ProductSLA_PLUS:
		return "Plus"
	case services.ProductSLA_BUSINESS:
		return "Business"
	case services.ProductSLA_ENTERPRISE:
		return "Enterprise"
	}

	return "Best Effort"
}

func (api *API) ProductLocations(locations []string) string {
	if len(locations) == 0 {
		return "unspecified"
	}

	for n, l := range locations {
		locations[n] = strings.Title(l)
	}

	return strings.Join(locations, ", ")
}

func (api *API) ProductLangs(langs []string) string {
	if len(langs) == 0 {
		return "unspecified"
	}

	for n, l := range langs {
		langs[n] = strings.Title(l)
	}

	return strings.Join(langs, ", ")
}

func (api *API) ProviderRating(rating uint32) string {
	var r string

	switch rating {
	case 0:
		r = colors.Black("* * * * *")
	case 1:
		r = fmt.Sprintf("%s %s", colors.Green("*"), colors.Black("* * * *"))
	case 2:
		r = fmt.Sprintf("%s %s", colors.Green("* *"), colors.Black("* * *"))
	case 3:
		r = fmt.Sprintf("%s %s", colors.Green("* * *"), colors.Black("* *"))
	case 4:
		r = fmt.Sprintf("%s %s", colors.Green("* * * *"), colors.Black("*"))
	case 5:
		r = colors.Green("* * * * *")
	default:
		r = colors.Black("n/a")
	}

	return fmt.Sprintf("%s%s%s", colors.DarkGreen("["), r, colors.DarkGreen("]"))
}

func productAvailable(s bool) string {
	if s {
		return output.StrEnabled("yes")
	}

	return output.StrDisabled("no")
}

func productClass(pc services.ProductClass) string {
	switch pc {
	case services.ProductClass_CLASS_CLOUD:
		return "Cloud"
	case services.ProductClass_CLASS_SERVICE:
		return "Professional Service"
	}

	return "-"
}

func productServiceCategory(sc services.ServiceCategory) string {
	switch sc {
	case services.ServiceCategory_SERVICE_ADVISORY_SERVICE:
		return "Advisory Service"
	case services.ServiceCategory_SERVICE_MANAGED_SERVICE:
		return "Managed Service"
	case services.ServiceCategory_SERVICE_INFRASTRUCTURE_PROJECT:
		return "Infrastructure Project"
	case services.ServiceCategory_SERVICE_AUTOMATION_PROJECT:
		return "Automation Project"
	case services.ServiceCategory_SERVICE_SOFTWARE_DEVELOPMENT:
		return "Software Development"
	case services.ServiceCategory_SERVICE_PROJECT_MANAGEMENT:
		return "Project Management"
	case services.ServiceCategory_SERVICE_SPECIAL_PROJECT:
		return "Special Project"
	case services.ServiceCategory_SERVICE_SPECIAL_TASK:
		return "Special Task"
	}

	return "-"
}

func productServiceType(st services.ServiceType) string {
	switch st {
	case services.ServiceType_SERVICE_TYPE_UNSPECIFIED:
		return "UNSPECIFIED"
	case services.ServiceType_SERVICE_ADVISORY_SERVICE_ARCHITECTURE:
		return "ADVISORY SERVICE ARCHITECTURE"
	case services.ServiceType_SERVICE_ADVISORY_SERVICE_SUPPORT:
		return "ADVISORY SERVICE SUPPORT"
	case services.ServiceType_SERVICE_MANAGED_SERVICE_CLOUD:
		return "MANAGED SERVICE CLOUD"
	case services.ServiceType_SERVICE_MANAGED_SERVICE_ONPREM:
		return "MANAGED SERVICE ONPREM"
	case services.ServiceType_SERVICE_MANAGED_SERVICE_INFRASTRUCTURE_MAINTENANCE:
		return "MANAGED SERVICE INFRASTRUCTURE MAINTENANCE"
	case services.ServiceType_SERVICE_MANAGED_SERVICE_SOFTWARE_MAINTENANCE:
		return "MANAGED SERVICE SOFTWARE MAINTENANCE"
	case services.ServiceType_SERVICE_INFRASTRUCTURE_PROJECT_NETWORK:
		return "INFRASTRUCTURE PROJECT NETWORK"
	case services.ServiceType_SERVICE_INFRASTRUCTURE_PROJECT_CLOUD:
		return "INFRASTRUCTURE PROJECT CLOUD"
	case services.ServiceType_SERVICE_INFRASTRUCTURE_PROJECT_ONPREM:
		return "INFRASTRUCTURE PROJECT ONPREM"
	case services.ServiceType_SERVICE_INFRASTRUCTURE_PROJECT_MIGRATION:
		return "INFRASTRUCTURE PROJECT MIGRATION"
	case services.ServiceType_SERVICE_AUTOMATION_PROJECT_SIMPLE:
		return "AUTOMATION PROJECT SIMPLE"
	case services.ServiceType_SERVICE_AUTOMATION_PROJECT_COMPLEX:
		return "AUTOMATION PROJECT COMPLEX"
	case services.ServiceType_SERVICE_SOFTWARE_DEVELOPMENT_SIMPLE:
		return "SOFTWARE DEVELOPMENT SIMPLE"
	case services.ServiceType_SERVICE_SOFTWARE_DEVELOPMENT_COMPLEX:
		return "SOFTWARE DEVELOPMENT COMPLEX"
	case services.ServiceType_SERVICE_SOFTWARE_DEVELOPMENT_SUPPORT:
		return "SOFTWARE DEVELOPMENT SUPPORT"
	case services.ServiceType_SERVICE_SOFTWARE_DEVELOPMENT_MAINTENANCE:
		return "SOFTWARE DEVELOPMENT MAINTENANCE"
	case services.ServiceType_SERVICE_SOFTWARE_DEVELOPMENT_FEATURE_REQUEST:
		return "SOFTWARE DEVELOPMENT FEATURE REQUEST"
	case services.ServiceType_SERVICE_PROJECT_MANAGEMENT_SIMPLE:
		return "PROJECT MANAGEMENT SIMPLE"
	case services.ServiceType_SERVICE_PROJECT_MANAGEMENT_COMPLEX:
		return "PROJECT MANAGEMENT COMPLEX"
	case services.ServiceType_SERVICE_SPECIAL_PROJECT_SIMPLE:
		return "SPECIAL PROJECT SIMPLE"
	case services.ServiceType_SERVICE_SPECIAL_PROJECT_COMPLEX:
		return "SPECIAL PROJECT COMPLEX"
	case services.ServiceType_SERVICE_SPECIAL_TASK_SIMPLE:
		return "SPECIAL TASK SIMPLE"
	case services.ServiceType_SERVICE_SPECIAL_TASK_COMPLEX:
		return "SPECIAL TASK COMPLEX"
	case services.ServiceType_SERVICE_SPECIAL_TASK_INCIDENT_RESPONSE:
		return "SPECIAL TASK INCIDENT RESPONSE"
	case services.ServiceType_SERVICE_SPECIAL_TASK_SECURITY_ASSESSMENT:
		return "SPECIAL TASK SECURITY ASSESSMENT"
	case services.ServiceType_SERVICE_SPECIAL_TASK_HEALTH_CHECKING:
		return "SPECIAL TASK HEALTH CHECKING"
	case services.ServiceType_SERVICE_SPECIAL_TASK_COMPLIANCE_ASSESSMENT:
		return "SPECIAL TASK COMPLIANCE ASSESSMENT"
	}

	return "-"
}

func productServiceScope(ss services.ServiceScope) string {
	switch ss {
	case services.ServiceScope_SCOPE_UNDEFINED:
		return "UNDEFINED"
	case services.ServiceScope_SCOPE_ALL:
		return "ALL"
	case services.ServiceScope_SCOPE_DESIGN:
		return "DESIGN"
	case services.ServiceScope_SCOPE_IMPLEMENTATION:
		return "IMPLEMENTATION"
	case services.ServiceScope_SCOPE_MAINTENANCE:
		return "MAINTENANCE"
	case services.ServiceScope_SCOPE_SUPPORT:
		return "SUPPORT"
	}

	return "-"
}
