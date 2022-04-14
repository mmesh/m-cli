package operation

func (api *API) List() {
	Output().List(operations())
}
