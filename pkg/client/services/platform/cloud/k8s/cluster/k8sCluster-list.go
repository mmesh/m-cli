package cluster

func (api *API) List() {
	Output().List(k8sClusters())
}
