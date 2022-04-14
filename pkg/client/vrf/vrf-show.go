package vrf

func (api *API) Show() {
	Output().Show(GetVRF(false))
}
