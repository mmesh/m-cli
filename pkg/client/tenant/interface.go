package tenant

import "mmesh.dev/m-cli/pkg/client/tenant/output"

type Interface interface {
	List()
	Show()
	New()
	Update()
	Delete()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
