package command

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-cli/pkg/auth"
	"mmesh.dev/m-cli/pkg/client/node"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-lib/pkg/mmid"
	"mmesh.dev/m-lib/pkg/mmp"
	"mmesh.dev/m-lib/pkg/mmp/cli"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) Exec(args []string) {
	n := node.GetNode(false)

	nxnc, grpcConn := grpc.GetNetworkAPIClient()
	defer grpcConn.Close()

	var execCommand string

	if len(args) == 0 {
		inputMsg := "Command:"
		helpMsg := "Command with args you want to execute on the target node"
		defaultCommand := "bash"
		c := input.GetInput(inputMsg, helpMsg, defaultCommand, survey.Required)

		args = strings.Split(c, " ")
	}

	execCommand = args[0]

	argc := len(args)
	if argc > 1 {
		args = args[1:argc]
	} else {
		args = []string{}
	}

	var wg sync.WaitGroup
	execWaitc := make(chan struct{}, 2)

	wg.Add(1)
	go grpc.Control(nxnc, &wg, execWaitc)

	dstID := mmid.NewMMNodeID(n.AccountID, n.TenantID, n.NetID, n.VRFID, n.NodeID)

	authKey, err := auth.GetAuthKey()
	if err != nil {
		msg.Error("Invalid or inexistent authorization key. Login to refresh your token.")
		os.Exit(1)
	}

	output.SectionHeader("mmesh remote management")

	mmp.NewShellRequest(authKey, dstID.String(), execCommand, args...)

	cli.Connected()

	output.CmdLog("Executing command via mmesh network..")
	fmt.Println()

	// go execClose(execWaitc)

	wg.Wait()
}

/*
func execClose(waitc chan struct{}) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-c
		switch sig {
		case os.Interrupt:
		case syscall.SIGINT:
		case syscall.SIGTERM:
		}

		msg.Info("Closing connection handlers...")
		waitc <- struct{}{}
	}()
}
*/
