package item

func (api *API) List() {
	Output().List(invoiceItems())
}
