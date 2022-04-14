package cluster

func (api *API) GetKubeConfig() {
	Output().KubeConfig(GetKubernetesCluster(false))
}
