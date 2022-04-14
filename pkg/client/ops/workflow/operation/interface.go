package operation

import "mmesh.dev/m-cli/pkg/client/ops/workflow/operation/output"

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
