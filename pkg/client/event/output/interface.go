package output

import "mmesh.dev/m-api-go/grpc/resources/events"

type Interface interface {
	ShowMetrics(em *events.EventMetrics)
	FailureProbability(em *events.EventMetrics) string
}
type API struct{}
