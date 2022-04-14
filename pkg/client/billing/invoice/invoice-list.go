package invoice

func (api *API) List() {
	Output().List(invoices())
}
