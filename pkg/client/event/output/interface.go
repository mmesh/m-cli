package output

import "mmesh.dev/m-api-go/grpc/resources/metrics"

type Interface interface {
	ShowMetrics(em *metrics.EventMetrics)
	FailureProbability(em *metrics.EventMetrics) string
}
type API struct{}
