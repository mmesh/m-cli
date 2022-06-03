package alert

import (
	"mmesh.dev/m-api-go/grpc/resources/events"
	"mmesh.dev/m-cli/pkg/client/alert/output"
)

type Interface interface {
	List()
	Show() *events.Alert
	Delete()
	NewNote(a *events.Alert)
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
