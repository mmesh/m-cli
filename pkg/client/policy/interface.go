package policy

import "mmesh.dev/m-cli/pkg/client/policy/output"

type Interface interface {
	Show()
	Import(yamlFile string)
	Export()
	SetDefault()
	SetRule()
	UnsetRule()
	Delete()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
