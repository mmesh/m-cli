package account

import (
	"context"
	"fmt"
	"os"

	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) Cancel() {
	a := fetchAccount()

	nxc, grpcConn := grpc.GetManagerProviderAPIClient(true)
	defer grpcConn.Close()

	confirmCancelation()

	s := output.Spinner()

	sr, err := nxc.CancelAccount(context.TODO(), a)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to cancel account")
	}

	s.Stop()

	output.Show(sr)
}

func confirmCancelation() {
	msg.Alert("You are about to cancel your mmesh account.")
	msg.Alert("All its remaining resources (workflows, users, etc) will be deleted.")
	msg.Alert("This action is irreversible, please, double check.")

	if !input.GetConfirm("Confirm account cancelation?", false) {
		fmt.Println()
		os.Exit(0)
	}

	if !input.GetConfirm("Last chance. Confirm account cancelation?", false) {
		fmt.Println()
		os.Exit(0)
	}
}
