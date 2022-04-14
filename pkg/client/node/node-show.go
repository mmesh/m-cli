package node

func (api *API) Show() {
	Output().Show(GetNode(false))
}
