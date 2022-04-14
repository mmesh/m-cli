package policy

import (
	"mmesh.dev/m-cli/pkg/client/vrf"
)

func (api *API) Show() {
	v := vrf.GetVRF(false)

	Output().Show(v.NetworkPolicy)
}
