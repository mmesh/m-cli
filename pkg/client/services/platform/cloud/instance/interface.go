package instance

import (
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-api-go/grpc/resources/services/catalog/cloud/compute"
	"mmesh.dev/m-cli/pkg/client/services/platform/cloud/instance/output"
)

type Interface interface {
	List(imgClass compute.ImageClass)
	Show(imgClass compute.ImageClass)
	Create(loginReq *auth.LoginRequest, imgClass compute.ImageClass)
	Delete(imgClass compute.ImageClass)
	PowerCycle(imgClass compute.ImageClass)
	PowerOn(imgClass compute.ImageClass)
	PowerOff(imgClass compute.ImageClass)
	Reboot(imgClass compute.ImageClass)
	Shutdown(imgClass compute.ImageClass)
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
