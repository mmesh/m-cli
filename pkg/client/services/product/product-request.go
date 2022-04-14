package product

import (
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/services"
	"mmesh.dev/m-cli/pkg/client/provider"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func RequestProductService(svcCategory services.ServiceCategory) *services.Product {
	var productOptID string
	var productsOpts []string
	productList := make(map[string]*services.Product)

	s := output.Spinner()

	providers := provider.ListProviders(services.ProviderType_PROFESSIONAL_SERVICES)

	for _, p := range providers {
		if p.Type != services.ProviderType_PROFESSIONAL_SERVICES {
			continue
		}

		// rating := getProviderRating(p.Rating)

		pl := products(p)

		if len(pl) == 0 {
			msg.Info("No objects found")
			os.Exit(1)
		}

		for _, product := range pl {
			if !product.Available {
				continue
			}

			if product.ServiceCategory != svcCategory {
				continue
			}

			price := Output().ProductPrice(product)
			if len(price) == 0 {
				continue
			}
			sla := Output().ProductSLA(product.SLA)

			productOptID = fmt.Sprintf(
				"Provider: %s | %s\n  Service: %s\n  Description: %s\n  SLA: %s\n  Price: %s\n  Langs: %s\n  Locations: %s\n",
				p.Name,
				p.WebsiteURL,
				// rating,
				product.Name,
				product.Description,
				sla,
				price,
				Output().ProductLangs(product.Langs),
				Output().ProductLocations(product.Locations),
			)
			productsOpts = append(productsOpts, productOptID)
			productList[productOptID] = product
		}
	}

	sort.Strings(productsOpts)

	s.Stop()

	if len(productsOpts) == 0 {
		msg.Alert("No service/products available at this moment")
		os.Exit(1)
	}

	productOptID = input.GetSelect("Product/Service:", "", productsOpts, survey.Required)

	return productList[productOptID]
}

func SelectServiceCategory() services.ServiceCategory {
	svcCategories := map[string]services.ServiceCategory{
		"Advisory Service":       services.ServiceCategory_SERVICE_ADVISORY_SERVICE,
		"Managed Service":        services.ServiceCategory_SERVICE_MANAGED_SERVICE,
		"Infrastructure Project": services.ServiceCategory_SERVICE_INFRASTRUCTURE_PROJECT,
		"Automation Project":     services.ServiceCategory_SERVICE_AUTOMATION_PROJECT,
		"Software Development":   services.ServiceCategory_SERVICE_SOFTWARE_DEVELOPMENT,
		"Project Management":     services.ServiceCategory_SERVICE_PROJECT_MANAGEMENT,
		"Special Project":        services.ServiceCategory_SERVICE_SPECIAL_PROJECT,
		"Special Task":           services.ServiceCategory_SERVICE_SPECIAL_TASK,
	}

	opts := make([]string, 0)

	for c := range svcCategories {
		opts = append(opts, c)
	}

	sort.Strings(opts)

	sc := input.GetSelect("Service Category:", "", opts, survey.Required)

	return svcCategories[sc]
}
