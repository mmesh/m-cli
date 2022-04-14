package rt

import "mmesh.dev/m-cli/pkg/client/rt/output"

type Interface interface {
	List()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
