package instance

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/services/catalog/cloud/compute"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) Delete(imgClass compute.ImageClass) {
	i := GetInstance(false, imgClass)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	msg.Alert("You are about to destroy a cloud server instance.")
	msg.Alert("All its resources (disk volumes, etc) will be destroyed too.")
	msg.Alert("This action is irreversible, please, double check.")

	output.ConfirmDeletion()

	s := output.Spinner()

	msg.Info("This action might take 1-5 minutes...")

	sr, err := nxc.DeleteCloudInstance(context.TODO(), i)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete cloud instance")
	}

	s.Stop()

	output.Show(sr)
}
