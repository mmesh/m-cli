package account

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/skratchdot/open-golang/open"
	"mmesh.dev/m-api-go/grpc/resources/account"
	"mmesh.dev/m-api-go/grpc/resources/billing"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) Subscription(a *account.Account, interactive bool) {
	if a == nil {
		a = fetchAccountStats()
	}

	Output().Service(a)

	if a.Service == nil {
		return
	}

	requiresPaymentMethod := false

	n := 0
	for _, s := range a.Service.Subscriptions {
		Output().Subscription(s, n)

		if !interactive {
			continue
		}

		switch s.LatestStripeInvoicePaymentIntentStatus {
		case "requires_payment_method":
			requiresPaymentMethod = true
		case "requires_action":
			if len(s.LatestStripeHostedInvoiceURL) > 0 {
				if input.GetConfirm(`Your bank/card issuer is requesting additional authentication
  to authorize an ongoing payment.

  Open the payment form now?`, true) {
					if err := open.Start(s.LatestStripeHostedInvoiceURL); err != nil {
						status.Error(err, "Unable to open URL in your browser")
					}

					fmt.Printf("\n%s %s\n\n", colors.DarkWhite("🢂"), colors.Black("Opening URL in your browser..."))
				} else {
					fmt.Println()
				}
			}
		}

		n++
	}

	if !interactive {
		return
	}

	if requiresPaymentMethod {
		if input.GetConfirm("A failed payment attempt requires your attention, open the Billing Portal now?", true) {
			api.BillingPortal(a)
		} else {
			fmt.Println()
		}
	} else {
		opt := !checkLimit(a)
		if !input.GetConfirm("Upgrade subscription now?", opt) {
			fmt.Println()
			return
		}

		api.BillingPortal(a)
	}
}

func (api *API) ApplyPromotion() {
	a := FetchAccount()

	api.Subscription(a, false)

	var sID string
	for _, s := range a.Service.Subscriptions {
		sID = s.StripeSubscriptionID
		break
	}

	apr := &billing.ApplyPromotionRequest{
		AccountID:            a.AccountID,
		StripeSubscriptionID: sID,
		PromotionCode:        input.GetInput("Promotion Code:", "", "", survey.Required),
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetBillingAPIClient(true)
	defer grpcConn.Close()

	sr, err := nxc.ApplyPromotion(context.TODO(), apr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to apply promotion code")
	}

	s.Stop()

	output.Show(sr)
}

func (api *API) BillingPortal(a *account.Account) {
	if a == nil {
		a = FetchAccount()
	}

	s := output.Spinner()

	nxc, grpcConn := grpc.GetBillingAPIClient(true)
	defer grpcConn.Close()

	customerPortalRequest := &billing.CustomerPortalRequest{
		AccountID:  a.AccountID,
		CustomerID: a.Owner.Customer.StripeCustomerID,
	}

	r, err := nxc.GetCustomerPortal(context.TODO(), customerPortalRequest)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get customer portal URL")
	}

	if err := open.Start(r.URL); err != nil {
		s.Stop()
		status.Error(err, "Unable to open URL in your browser")
	}

	s.Stop()

	fmt.Printf("\n%s %s\n\n", colors.DarkWhite("🢂"), colors.Black("Opening Billing Portal in your browser..."))
}

func checkLimit(a *account.Account) bool {
	if a.Usage == nil {
		return true
	}

	if a.Usage.Limit == nil {
		return true
	}

	if !a.Usage.Limit.OverLimit {
		return true
	}

	msg.Warn("Your Free account is over limits, upgrade recommended.")

	return false
}
