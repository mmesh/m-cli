package vrf

import "mmesh.dev/m-cli/pkg/client/vrf/output"

type Interface interface {
	List()
	Show()
	New()
	Update()
	Delete()
	DeleteIPAMEntry()
	// EnableVRFRelayService()
	// DisableVRFRelayService()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
