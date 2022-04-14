package output

import (
	"fmt"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/services"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(p *services.Product) {
	output.SectionHeader("Service: Product Details")
	output.TitleT1("Product Information")

	t := table.New()

	t.BulkData([][]string{
		{colors.Black("Service ID"), colors.DarkWhite(p.ServiceID)},
		{colors.Black("Provider ID"), colors.DarkWhite(p.ProviderID)},
		{colors.Black("Product ID"), colors.DarkWhite(p.ProductID)},
		{colors.Black("Name"), colors.DarkWhite(p.Name)},
		{colors.Black("Description"), colors.DarkWhite(p.Description)},
		{colors.Black("Available"), productAvailable(p.Available)},
		{colors.Black("Price"), colors.DarkWhite(api.ProductPrice(p))},
		{colors.Black("SLA"), colors.DarkWhite(api.ProductSLA(p.SLA))},
		{colors.Black("Class"), colors.DarkWhite(productClass(p.Class))},
		{colors.Black("Category"), colors.DarkWhite(productServiceCategory(p.ServiceCategory))},
		{colors.Black("Type"), output.StrFixedField(productServiceType(p.ServiceType))},
		{colors.Black("Scope"), output.StrFixedField(productServiceScope(p.ServiceScope))},
		{colors.Black("Locations"), colors.DarkWhite(api.ProductLocations(p.Locations))},
		{colors.Black("Langs"), colors.DarkWhite(api.ProductLangs(p.Langs))},
		{colors.Black("Creation Date"), colors.DarkWhite(time.Unix(p.CreationDate, 0).String())},
		{colors.Black("Last Modified"), colors.DarkWhite(time.Unix(p.LastModified, 0).String())},
	})

	t.Render()

	fmt.Println()
}
