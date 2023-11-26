package node

import "mmesh.dev/m-cli/pkg/client/node/output"

type Interface interface {
	AddNode()
	// GetInstallationWebhook()
	ListByTenant()
	ListBySubnet()
	Show()
	Delete()
	Connect()
	Disconnect()
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
