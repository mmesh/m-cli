package node

func (api *API) List() {
	//output.Show(nl)
	Output().List(nodes())
}
