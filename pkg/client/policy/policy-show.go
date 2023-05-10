package policy

import (
	"mmesh.dev/m-cli/pkg/client/subnet"
)

func (api *API) Show() {
	s := subnet.GetSubnet(false)

	Output().Show(s)
}
