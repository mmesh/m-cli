package account

import (
	"mmesh.dev/m-api-go/grpc/resources/account"
	"mmesh.dev/m-cli/pkg/client/account/output"
)

type Interface interface {
	Show()
	// Settings()
	Cancel()
	Subscription(a *account.Account, interactive bool)
	// ApplyPromotion()
	BillingPortal(a *account.Account)
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
