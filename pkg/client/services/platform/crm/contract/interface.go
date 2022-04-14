package contract

import (
	"mmesh.dev/m-cli/pkg/client/services/platform/crm/contract/output"
)

type Interface interface {
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
