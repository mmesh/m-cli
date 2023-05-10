package output

import (
	"fmt"
	"sort"

	"mmesh.dev/m-api-go/grpc/resources/billing"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(invoiceItems map[string]*billing.InvoiceItem) {
	output.SectionHeader("Billing: Billable Items")
	output.TitleT1("Billable Item List")

	t := table.New()
	t.Header(colors.Black("RESOURCE"), colors.Black("CREATION DATE"))

	biItems := make([]*billing.InvoiceItem, 0)

	for _, ii := range invoiceItems {
		biItems = append(biItems, ii)
	}

	sort.SliceStable(biItems, func(i, j int) bool {
		return biItems[i].CreationDate < biItems[j].CreationDate
	})

	for _, ii := range biItems {
		tm := output.Datetime(ii.CreationDate)
		t.AddRow(ii.ShortDescription, colors.DarkWhite(tm))
	}

	t.Render()
	fmt.Println()
}
