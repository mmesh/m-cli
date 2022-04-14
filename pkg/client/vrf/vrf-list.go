package vrf

func (api *API) List() {
	Output().List(vrfs())
}
