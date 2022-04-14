package event

import "mmesh.dev/m-cli/pkg/client/event/output"

func Output() output.Interface {
	return &output.API{}
}
