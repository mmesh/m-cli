package output

import (
	"fmt"
	"strings"

	"mmesh.dev/m-api-go/grpc/resources/services/platform/crm"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(contracts map[string]*crm.Contract) {
	output.SectionHeader("Services: Contracts")
	output.TitleT1("Service Contract List")

	t := table.New()
	t.Header(colors.Black("SERVICE PROVIDER"), colors.Black("CONTRACT ID"), colors.Black("PRODUCT / SERVICE"))

	for _, sc := range contracts {
		providerID := output.StrNormal(strings.Title(sc.Spec.Product.ProviderID))
		scID := colors.DarkWhite(sc.ContractID)
		t.AddRow(providerID, scID, sc.Spec.Product.Name)
	}

	t.Render()
	fmt.Println()
}
