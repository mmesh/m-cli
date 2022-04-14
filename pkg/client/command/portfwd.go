package command

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-cli/pkg/auth"
	"mmesh.dev/m-cli/pkg/client/node"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/mmp"
	"mmesh.dev/m-lib/pkg/mmp/cli"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

const proto string = "TCP"

func (api *API) PortFwd() {
	n := node.GetNode(false)

	nxnc, grpcConn := grpc.GetNetworkAPIClient()
	defer grpcConn.Close()

	var lp, rp string
	helpText := "Remote TCP port on target node"
	prompt := &survey.Input{Message: "Remote TCP Port:", Help: helpText}
	if err := survey.AskOne(prompt, &rp, survey.WithValidator(input.ValidPort), survey.WithIcons(input.SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}

	helpText = "Local TCP port"
	prompt = &survey.Input{Message: "Local TCP Port:", Help: helpText}
	if err := survey.AskOne(prompt, &lp, survey.WithValidator(input.ValidPort), survey.WithIcons(input.SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}

	srcP, err := strconv.Atoi(lp)
	if err != nil {
		status.Error(err, "Invalid local port")
	}
	dstP, err := strconv.Atoi(rp)
	if err != nil {
		status.Error(err, "Invalid destination port")
	}

	srcID := viper.GetString("mm.id")
	dstID := fmt.Sprintf("%s:%s:%s:%s:%s", n.AccountID, n.TenantID, n.NetID, n.VRFID, n.NodeID)

	var wg sync.WaitGroup
	portFwdWaitc := make(chan struct{}, 2)

	wg.Add(1)
	go grpc.Control(nxnc, &wg, portFwdWaitc)

	authKey, err := auth.GetAuthKey()
	if err != nil {
		msg.Alert("Invalid or inexistent authorization key. Login to refresh your token.")
		os.Exit(1)
	}

	output.SectionHeader("mmesh remote management")

	mmp.NewPortFwd(authKey, uint32(srcP), uint32(dstP), proto, srcID, dstID, true)

	cli.Connected()

	output.CmdLog(fmt.Sprintf("Activating mmesh port fowarding to node %s (%s/%s)", n.NodeID, proto, rp))

	go pfClose(portFwdWaitc, srcP, dstP, proto, srcID, dstID)

	wg.Wait()

	cli.Disconnected()
}

func pfClose(waitc chan struct{}, srcP, dstP int, proto, srcID, dstID string) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		var wg sync.WaitGroup
		<-c
		fmt.Println()
		output.CmdLog("Closing connections...")

		wg.Add(1)
		mmp.ClosePortFwd(uint32(srcP), uint32(dstP), proto, srcID, dstID, &wg)

		wg.Wait()
		waitc <- struct{}{}
	}()
}
