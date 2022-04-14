package product

func (api *API) Show(publicCatalog bool) {
	Output().Show(GetProduct(publicCatalog))
}
