package node

import "mmesh.dev/m-cli/pkg/client/node/output"

type Interface interface {
	AddNode()
	// GetInstallationWebhook()
	List()
	Show()
	Delete()
	Metrics()
	// ResetNetworkTraffic()
	ShowEndpoint()
	DeleteEndpoint()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
