package opportunity

import (
	"mmesh.dev/m-cli/pkg/client/services/platform/crm/opportunity/output"
)

type Interface interface {
	List()
	Show()
	Delete()
	SetMilestone()
	SetOutcome()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
