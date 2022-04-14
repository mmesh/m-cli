package output

import (
	"fmt"
	"sort"
	"time"

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

	tmSort := make([]string, 0)
	sortedItems := make(map[string]*billing.InvoiceItem, 0)
	for _, ii := range invoiceItems {
		tm := time.Unix(ii.CreationDate, 0).String()
		tmSort = append(tmSort, tm)
		sortedItems[tm] = ii
	}
	sort.Strings(tmSort)

	for _, tm := range tmSort {
		ii := sortedItems[tm]
		tm := output.Datetime(ii.CreationDate)
		t.AddRow(ii.ShortDescription, colors.DarkWhite(tm))
	}

	t.Render()
	fmt.Println()
}
