package alert

import "mmesh.dev/m-cli/pkg/client/alert/output"

type Interface interface {
	List()
	Show()
	Delete()
	NewComment()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
