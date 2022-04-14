package instance

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/services/catalog/cloud/compute"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) PowerCycle(imgClass compute.ImageClass) {
	i := GetInstance(false, imgClass)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	s := output.Spinner()

	i, err := nxc.PowerCycleCloudInstance(context.TODO(), i)
	if err != nil {
		status.Error(err, "Unable to power cycle the cloud instance")
	}

	s.Stop()

	Output().Show(i)
}

func (api *API) PowerOn(imgClass compute.ImageClass) {
	i := GetInstance(false, imgClass)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	s := output.Spinner()

	i, err := nxc.PowerOnCloudInstance(context.TODO(), i)
	if err != nil {
		status.Error(err, "Unable to power-on the cloud instance")
	}

	s.Stop()

	Output().Show(i)
}

func (api *API) PowerOff(imgClass compute.ImageClass) {
	i := GetInstance(false, imgClass)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	s := output.Spinner()

	i, err := nxc.PowerOffCloudInstance(context.TODO(), i)
	if err != nil {
		status.Error(err, "Unable to power-off the cloud instance")
	}

	s.Stop()

	Output().Show(i)
}

func (api *API) Reboot(imgClass compute.ImageClass) {
	i := GetInstance(false, imgClass)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	s := output.Spinner()

	i, err := nxc.RebootCloudInstance(context.TODO(), i)
	if err != nil {
		status.Error(err, "Unable to reboot the cloud instance")
	}

	s.Stop()

	Output().Show(i)
}

func (api *API) Shutdown(imgClass compute.ImageClass) {
	i := GetInstance(false, imgClass)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	s := output.Spinner()

	i, err := nxc.ShutdownCloudInstance(context.TODO(), i)
	if err != nil {
		status.Error(err, "Unable to shutdown the cloud instance")
	}

	s.Stop()

	Output().Show(i)
}
