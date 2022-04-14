package instance

import "mmesh.dev/m-api-go/grpc/resources/services/catalog/cloud/compute"

func (api *API) Show(imgClass compute.ImageClass) {
	Output().Show(GetInstance(false, imgClass))
}
