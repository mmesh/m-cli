package item

func (api *API) Show() {
	Output().Show(GetBilledItem())
}
