package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/services"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) List(products map[string]*services.Product) {
	output.SectionHeader("Service: Products")
	output.TitleT1("Product List")

	if len(products) == 0 {
		msg.Info("No records found")
		return
	}

	t := table.New()

	t.SetColWidth(36)

	t.Header(colors.Black("PRODUCT"), colors.Black("DESCRIPTION"))

	t.SetAutoWrapText(true)
	t.SetReflowDuringAutoWrap(true)
	// t.SetRowLine("─")
	t.SetRowLine("-")

	for _, p := range products {
		pName := fmt.Sprintf("%s", colors.DarkWhite(p.Name))
		t.AddRow(pName, p.Description)
	}

	t.Render()
	fmt.Println()
}
