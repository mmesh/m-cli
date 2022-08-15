package invoice

import (
	"fmt"

	"github.com/skratchdot/open-golang/open"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show() {
	i := GetInvoice()

	Output().Show(i)

	switch i.PaymentIntentStatus {
	case "requires_payment_method":
		if input.GetConfirm("A failed payment attempt requires your attention. Open the Billing Portal now?", true) {
			account.Resource().BillingPortal(nil)
		} else {
			fmt.Println()
		}
	case "requires_action":
		if len(i.HostedInvoiceURL) > 0 {
			if input.GetConfirm(`Your bank/card issuer is requesting additional authentication
  to authorize this payment.

  Open the payment form now?`, true) {
				if err := open.Start(i.HostedInvoiceURL); err != nil {
					status.Error(err, "Unable to open URL in your browser")
				}

				fmt.Printf("\n%s %s\n\n", colors.DarkWhite("🢂"), colors.Black("Opening payment form in your browser..."))
			} else {
				fmt.Println()
			}
		}
	}

	if i.Status == "paid" {
		if len(i.HostedInvoiceURL) > 0 {
			if input.GetConfirm("Download invoice?", false) {
				if err := open.Start(i.HostedInvoiceURL); err != nil {
					status.Error(err, "Unable to open URL in your browser")
				}

				fmt.Printf("\n%s %s\n\n", colors.DarkWhite("🢂"), colors.Black("Opening invoice in your browser..."))
			} else {
				fmt.Println()
			}
		}
	}
}
