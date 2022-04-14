package item

/*
import (
	"context"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Delete() {
	bi := GetBilledItem()

	nxc, grpcConn := grpc.GetBillingAPIClient(true)
	defer grpcConn.Close()

	output.ConfirmDeletion()

	s := output.Spinner()

	sr, err := nxc.DeleteBilledItem(context.TODO(), bi)
	if err != nil {
		status.Error(err, "Unable to delete billed-item")
	}

	s.Stop()

	output.Show(sr)
}
*/
