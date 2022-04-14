package invoice

import "mmesh.dev/m-cli/pkg/client/billing/invoice/output"

type Interface interface {
	List()
	Show()
	Delete()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
