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

func (api *API) List(invoices map[string]*billing.Invoice) {
	output.SectionHeader("Billing: Invoices")
	output.TitleT1("Invoice List")

	t := table.New()
	t.Header(colors.Black("PERIOD"), colors.Black("INVOICE #"), colors.Black("STATUS"))

	tmSort := make([]string, 0)
	sortedInvoices := make(map[string]*billing.Invoice, 0)
	for _, i := range invoices {
		tm := time.Unix(i.CreationDate, 0).String()
		tmSort = append(tmSort, tm)
		sortedInvoices[tm] = i
	}
	sort.Strings(tmSort)

	for _, tm := range tmSort {
		i := sortedInvoices[tm]
		status := invoiceStatus(i.Status)
		t.AddRow(colors.DarkWhite(i.Period), i.InvoiceNumber, status)
	}

	t.Render()
	fmt.Println()
}
