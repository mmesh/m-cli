package item

import (
	"context"
	"fmt"
	"os"
	"sort"

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

func GetBilledItem() *billing.InvoiceItem {
	iil := invoiceItems()

	if len(iil) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var invoiceItemOptID string
	invoiceItemsOpts := make([]string, 0)
	invoiceItems := make(map[string]*billing.InvoiceItem)

	for _, ii := range iil {
		var status string
		if ii.CancelationDate > 0 {
			status = "[CANCELED]"
		}
		creationDate := output.Datetime(ii.CreationDate)
		invoiceItemOptID = fmt.Sprintf("%s │ %s %s", creationDate, ii.ShortDescription, status)
		invoiceItemsOpts = append(invoiceItemsOpts, invoiceItemOptID)
		invoiceItems[invoiceItemOptID] = ii
	}

	sort.Strings(invoiceItemsOpts)

	invoiceItemOptID = input.GetSelect("Item:", "", invoiceItemsOpts, survey.Required)

	vars.BilledItemID = invoiceItems[invoiceItemOptID].ItemID

	return invoiceItems[invoiceItemOptID]
}

func invoiceItems() map[string]*billing.InvoiceItem {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetBillingAPIClient(true)
	defer grpcConn.Close()

	lr := &billing.ListInvoiceItemsRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	invoiceItems := make(map[string]*billing.InvoiceItem)

	for {
		bil, err := nxc.ListBilledItems(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list billed items")
		}

		for _, bi := range bil.InvoiceItems {
			invoiceItems[bi.ItemID] = bi
		}

		if len(bil.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = bil.Meta.NextPageToken
		} else {
			break
		}
	}

	return invoiceItems
}
