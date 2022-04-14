package output

import "mmesh.dev/m-api-go/grpc/resources/events"

type Interface interface {
	List(alerts map[string]*events.Alert)
	Show(a *events.Alert)

	AlertStatus(s string) string
}
type API struct{}
