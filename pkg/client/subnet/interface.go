package subnet

import "mmesh.dev/m-cli/pkg/client/subnet/output"

type Interface interface {
	List()
	Show()
	New()
	Update()
	Delete()
	DeleteIPAMEntry()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
