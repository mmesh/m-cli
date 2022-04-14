package cluster

import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	kc := GetKubernetesCluster(false)

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	sr, err := nxc.DeleteCloudKubernetesCluster(context.TODO(), kc)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to delete cloud kubernetes cluster")
	}

	s.Stop()

	output.Show(sr)
}
