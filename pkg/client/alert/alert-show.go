package alert

import "mmesh.dev/m-api-go/grpc/resources/events"

func (api *API) Show() *events.Alert {
	a := getAlert()

	Output().Show(a)

	return a
}
