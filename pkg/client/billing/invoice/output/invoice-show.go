package output

import (
	"fmt"
	"strings"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/billing"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(i *billing.Invoice) {
	// fmt.Println() // extra line needed due to the long select option above
	output.SectionHeader("Billing: Invoice Details")
	output.TitleT1("Invoice Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(i.AccountID))
	t.AddRow(colors.Black("Customer"), colors.DarkWhite(i.CustomerEmail))
	// t.AddRow(colors.Black("Invoice ID"), colors.DarkWhite(i.InvoiceID))

	if len(i.Description) > 0 {
		t.AddRow(colors.Black("Description"), colors.DarkWhite(i.Description))
	}

	t.AddRow(colors.Black("Invoice Number"), colors.DarkWhite(i.InvoiceNumber))

	if len(i.ReceiptNumber) > 0 {
		t.AddRow(colors.Black("Receipt Number"), colors.DarkWhite(i.ReceiptNumber))
	}
	if len(i.StatementDescriptor) > 0 {
		t.AddRow(colors.Black("Statement Descriptor"), colors.DarkWhite(i.StatementDescriptor))
	}

	// t.AddRow(colors.Black("Period"), colors.DarkWhite(i.Period))
	tm := time.Unix(i.PeriodStart, 0).String()
	t.AddRow(colors.Black("Period Start"), colors.DarkWhite(tm))
	tm = time.Unix(i.PeriodEnd, 0).String()
	t.AddRow(colors.Black("Period End"), colors.DarkWhite(tm))

	currency := colors.Yellow(strings.ToUpper(i.Currency))
	currency = fmt.Sprintf("%s%s%s", colors.DarkYellow("["), currency, colors.DarkYellow("]"))
	t.AddRow(colors.Black("Currency"), currency)
	t.AddRow("", "")
	t.AddRow(colors.Black("Starting Balance"), output.AmountMoney(i.StartingBalance, i.Currency))
	t.AddRow(colors.Black("Ending Balance"), output.AmountMoney(i.EndingBalance, i.Currency))

	t.AddRow("", "")
	invItems := fmt.Sprintf("%d", len(i.SummaryItems))
	t.AddRow(colors.Black("Invoice Items"), colors.DarkWhite(invItems))

	t.AddRow("", "")
	t.AddRow(colors.Black("Subtotal"), output.AmountMoney(i.Subtotal, i.Currency))
	if i.TotalDiscount > 0 {
		t.AddRow(colors.Black("Discount"), output.AmountMoney(-i.TotalDiscount, i.Currency))
	}
	t.AddRow(colors.Black("Taxes"), output.AmountMoney(i.Tax, i.Currency))
	t.AddRow(colors.Black("Total"), output.AmountMoney(i.Total, i.Currency))

	t.AddRow("", "")
	t.AddRow(colors.Black("Amount Due"), output.AmountMoney(i.AmountDue, i.Currency))
	t.AddRow(colors.Black("Amount Paid"), output.AmountMoney(i.AmountPaid, i.Currency))
	t.AddRow(colors.Black("Amount Remaining"), output.AmountMoney(i.AmountRemaining, i.Currency))

	if i.DueDate != 0 {
		if !i.Paid {
			daysUntilDue := time.Now().Sub(time.Unix(i.DueDate, 0)).Hours() / 24

			var daysLeft string
			if daysUntilDue < 0 {
				daysLeft = colors.Red(fmt.Sprintf("%.0f", daysUntilDue))
			} else {
				daysLeft = colors.Blue(fmt.Sprintf("%.0f", daysUntilDue))
			}
			t.AddRow(colors.Black("Days until Due"), daysLeft)
		}
		tm := time.Unix(i.DueDate, 0).String()
		t.AddRow(colors.Black("Due Date"), colors.DarkWhite(tm))
	}

	t.AddRow("", "")
	t.AddRow(colors.Black("Payment Attempts"), colors.DarkWhite(fmt.Sprintf("%d", i.AttemptCount)))

	if len(i.PaymentIntentStatus) > 0 {
		t.AddRow(colors.Black("Payment Status"), paymentStatus(i.PaymentIntentStatus))
	}

	if i.NextPaymentAttempt != 0 {
		tm := time.Unix(i.NextPaymentAttempt, 0).String()
		t.AddRow(colors.Black("Next Payment Attempt"), colors.DarkWhite(tm))
	}

	t.AddRow("", "")
	t.AddRow(colors.Black("Invoice Status"), invoiceStatus(i.Status))

	t.AddRow("", "")
	tm = time.Unix(i.CreationDate, 0).String()
	t.AddRow(colors.Black("Created"), colors.DarkWhite(tm))
	tm = time.Unix(i.LastModified, 0).String()
	t.AddRow(colors.Black("Last Modified"), colors.DarkWhite(tm))

	t.Render()
	fmt.Println()

	// if len(i.InvoicePDF) > 0 {
	// 	// fmt.Printf("Invoice PDF: %s\n\n", colors.Black(i.InvoicePDF))
	// 	if err := open.Start(i.InvoicePDF); err != nil {
	// 		status.Error(err, "Unable to open URL in your browser")
	// 	}
	// 	fmt.Printf("\n%s %s\n\n", colors.DarkWhite("🢂"), colors.Black("Opening invoice PDF in your browser..."))
	// }

	// invoice summary items
	if len(i.SummaryItems) == 0 {
		return
	}

	if input.GetConfirm("Show invoice items?", true) {
		fmt.Println()

		output.SubTitleT2("Invoice Items")

		n := 0
		for _, ii := range i.SummaryItems {
			n++
			showItem(ii, i.Currency, n, true)
		}
	}

	// invoice detail items
	if len(i.Items) == 0 {
		return
	}

	if input.GetConfirm("Show billed items details?", false) {
		fmt.Println()

		output.SubTitleT2("Billed Items")

		n := 0
		for _, ii := range i.Items {
			n++
			showItem(ii, i.Currency, n, false)
		}
	}
}

func paymentStatus(s string) string {
	ss := strings.ToLower(s)

	switch s {
	case "requires_payment_method":
		return output.StrError(ss)
	case "requires_confirmation":
		return output.StrNormal(ss)
	case "requires_action":
		return output.StrWarning(ss)
	case "processing":
		return output.StrNormal(ss)
	case "requires_capture":
		return output.StrInactive(ss)
	case "canceled":
		return output.StrInactive(ss)
	case "succeeded":
		return output.StrOk(ss)
	}

	return output.StrNormal(ss)
}

func invoiceStatus(s string) string {
	ss := strings.ToUpper(s)

	switch s {
	case "draft":
		return output.StrNormal(ss)
	case "open":
		return output.StrNormal(ss)
	case "paid":
		return output.StrOk(ss)
	case "uncollectible":
		return output.StrError(ss)
	case "void":
		return output.StrInactive(ss)
	}

	return output.StrNormal(ss)
}

func showItem(ii *billing.InvoiceItem, currency string, n int, isSummary bool) {
	t := table.New()

	svcID := fmt.Sprintf("%s%s%s %s", colors.DarkMagenta("["), colors.Magenta(fmt.Sprintf("%003d", n)), colors.DarkMagenta("]"), colors.Black("Service ID"))
	t.AddRow(svcID, colors.DarkWhite(ii.ServiceID))
	t.AddRow(colors.Black("      Provider ID"), colors.DarkWhite(ii.ProviderID))

	if !isSummary {
		if len(ii.ItemID) > 0 {
			t.AddRow(colors.Black("      Item ID"), colors.DarkWhite(ii.ItemID))
		}
	}

	if ii.StartTime > 0 {
		tm := time.Unix(ii.StartTime, 0).String()
		t.AddRow(colors.Black("      Start Time"), colors.DarkWhite(tm))
	}
	if ii.EndTime > 0 {
		tm := time.Unix(ii.EndTime, 0).String()
		t.AddRow(colors.Black("      End Time"), colors.DarkWhite(tm))
	}

	if len(ii.Category) > 0 {
		t.AddRow(colors.Black("      Category"), colors.DarkWhite(ii.Category))
	}

	if !isSummary {
		if len(ii.ProductName) > 0 {
			t.AddRow(colors.Black("      Product Name"), colors.DarkWhite(ii.ProductName))
		}
	}

	if len(ii.ShortDescription) > 0 {
		t.AddRow(colors.Black("      Description"), colors.DarkWhite(ii.ShortDescription), "")
	} else {
		t.AddRow(colors.Black("      Description"), colors.DarkWhite(ii.LongDescription), "")
	}

	if !isSummary {
		if len(ii.UnitPrice) > 0 {
			t.AddRow(colors.Black("      Monthly Price"), colors.DarkWhite(ii.UnitPrice))
		}
		if len(ii.Discount) > 0 {
			t.AddRow(colors.Black("      Discount"), colors.DarkWhite(ii.Discount))
		}
		if len(ii.PriceWithDiscount) > 0 {
			t.AddRow(colors.Black("      Price w/Discount"), colors.DarkWhite(ii.PriceWithDiscount))
		}
	}

	t.AddRow(colors.Black("      Amount"), output.AmountMoney(ii.Amount, currency))

	t.Render()
	fmt.Println()
}
