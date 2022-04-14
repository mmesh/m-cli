package sg

func (api *API) Show() {
	Output().Show(GetSecurityGroup(false))
}
