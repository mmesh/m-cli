package instance

import (
	"mmesh.dev/m-api-go/grpc/resources/services/catalog/cloud/compute"
)

func (api *API) List(imgClass compute.ImageClass) {
	Output().List(instances(imgClass))
}
