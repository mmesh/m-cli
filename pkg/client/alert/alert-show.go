package alert

func (api *API) Show() {
	Output().Show(GetAlert())
}
