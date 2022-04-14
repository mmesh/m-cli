package invoice

func (api *API) Show() {
	Output().Show(GetInvoice())
}
