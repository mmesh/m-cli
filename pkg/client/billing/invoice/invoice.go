package invoice

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/billing"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func GetInvoice() *billing.Invoice {
	il := invoices()

	if len(il) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var invoiceOptID string
	invoicesOpts := make([]string, 0)
	invoices := make(map[string]*billing.Invoice)

	for _, i := range il {
		var invoiceNumber string
		if len(i.InvoiceNumber) == 0 {
			invoiceNumber = "n/a"
		} else {
			invoiceNumber = i.InvoiceNumber
		}
		invoiceOptID = fmt.Sprintf("%s │ %s [%s]", i.Period, invoiceNumber, strings.ToUpper(i.Status))
		invoicesOpts = append(invoicesOpts, invoiceOptID)
		invoices[invoiceOptID] = i
	}

	sort.Strings(invoicesOpts)

	invoiceOptID = input.GetSelect("Invoice:", "", invoicesOpts, survey.Required)

	vars.InvoiceID = invoices[invoiceOptID].InvoiceID

	return invoices[invoiceOptID]
}

func invoices() map[string]*billing.Invoice {
	a := account.GetAccount()

	s := output.Spinner()

	nxc, grpcConn := grpc.GetBillingAPIClient(true)
	defer grpcConn.Close()

	lr := &billing.ListInvoicesRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	invoices := make(map[string]*billing.Invoice)

	for {
		il, err := nxc.ListInvoices(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list invoices")
		}

		for _, i := range il.Invoices {
			invoices[i.InvoiceID] = i
		}

		if len(il.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = il.Meta.NextPageToken
		} else {
			break
		}
	}

	s.Stop()

	return invoices
}
