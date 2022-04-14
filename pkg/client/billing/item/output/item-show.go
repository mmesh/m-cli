package output

import (
	"fmt"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/billing"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(bi *billing.InvoiceItem) {
	output.SectionHeader("Billing: Billable Item Details")
	output.TitleT1("Billable Item Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(bi.AccountID))
	t.AddRow(colors.Black("Service ID"), colors.DarkWhite(bi.ServiceID))
	// t.AddRow(colors.Black("Provider ID"), colors.DarkWhite(strings.Title(bi.ProviderID)))
	t.AddRow(colors.Black("Provider ID"), colors.DarkWhite(bi.ProviderID))
	t.AddRow(colors.Black("Item ID"), colors.DarkWhite(bi.ItemID))

	period := "-"
	if len(bi.Period) > 0 {
		period = bi.Period
	}
	t.AddRow(colors.Black("Last Billed Period"), colors.DarkWhite(period))

	if bi.StartTime > 0 {
		tm := time.Unix(bi.StartTime, 0).String()
		t.AddRow(colors.Black("Start Time"), colors.DarkWhite(tm))
	}
	if bi.EndTime > 0 {
		tm := time.Unix(bi.EndTime, 0).String()
		t.AddRow(colors.Black("End Time"), colors.DarkWhite(tm))
	}

	t.AddRow(colors.Black("Category"), colors.DarkWhite(bi.Category))
	t.AddRow(colors.Black("Product Name"), colors.DarkWhite(bi.ProductName))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(bi.ShortDescription), "")

	// t.AddRow(colors.Black("Item Details"), colors.DarkWhite(bi.LongDescription), "")

	// if len(bi.Currency) == 0 {
	// 	bi.Currency = "usd"
	// }

	// c := colors.Yellow(strings.ToUpper(bi.Currency))
	// c = fmt.Sprintf("%s%s%s", colors.DarkYellow("["), c, colors.DarkYellow("]"))
	// t.AddRow(colors.Black("Currency"), c)

	// if bi.UnitPrice > 0 {
	// 	discountText := colors.DarkYellow(" (discounts NOT included)")
	// 	t.AddRow(colors.Black("Unit Price"), output.AmountMoney(bi.UnitPrice, bi.Currency)+discountText)
	// }

	if bi.Preview {
		t.AddRow(colors.Black("Status"), output.StrNormal("preview"))
	}
	// t.AddRow("Amount"), output.AmountMoney(bi.Amount, currency))

	// fmt.Printf("\nItem Details:\n")
	// fmt.Println(colors.DarkCyan(bi.LongDescription))

	if bi.CreationDate > 0 {
		tm := time.Unix(bi.CreationDate, 0).String()
		t.AddRow(colors.Black("Creation Date"), colors.DarkWhite(tm))
	}
	if bi.CancelationDate > 0 {
		tm := time.Unix(bi.CancelationDate, 0).String()
		t.AddRow(colors.Black("Cancelation Date"), colors.DarkWhite(tm))
	}

	if bi.LastModified > 0 {
		tm := time.Unix(bi.LastModified, 0)
		t.AddRow(colors.Black("Last Modified"), colors.DarkWhite(tm.String()))
	}

	t.Render()
	fmt.Println()
}
