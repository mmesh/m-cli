package cluster

func (api *API) Show() {
	Output().Show(GetKubernetesCluster(false))
}
