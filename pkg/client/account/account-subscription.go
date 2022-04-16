package account

import (
	"context"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/skratchdot/open-golang/open"
	"mmesh.dev/m-api-go/grpc/resources/billing"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Subscription() {
	a := FetchAccount()

	Output().Subscriptions(a)

	if !input.GetConfirm("Update subscription now?", false) {
		fmt.Println()
		return
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

	fmt.Printf("\n%s %s\n\n", colors.DarkWhite("🢂"), colors.Black("Opening Billing Portal URL in your browser..."))
}

func (api *API) ApplyPromotion() {
	a := FetchAccount()

	Output().Subscriptions(a)

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
