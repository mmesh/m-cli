package opportunity

func (api *API) List() {
	//output.Show(scl)
	Output().List(opportunities())
}
