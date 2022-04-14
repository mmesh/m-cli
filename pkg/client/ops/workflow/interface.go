package workflow

import "mmesh.dev/m-cli/pkg/client/ops/workflow/output"

type Interface interface {
	List()
	Show()
	Set(yamlFile string)
	Delete()
	Enable()
	Disable()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
