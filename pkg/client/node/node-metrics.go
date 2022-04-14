package node

func (api *API) Metrics() {
	Output().Metrics(GetNode(false))
}
