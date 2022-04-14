package account

import "mmesh.dev/m-cli/pkg/client/account/output"

type Interface interface {
	New()
	Show()
	Stats()
	Settings()
	Cancel()
	Subscription()
	ApplyPromotion()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
