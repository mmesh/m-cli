package command

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/auth"
	"mmesh.dev/m-cli/pkg/client/node"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/mm"
	"mmesh.dev/m-lib/pkg/mmp/stream/protos/portfwd"
	"mmesh.dev/m-lib/pkg/mmp/stream/utils/cli"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

const proto string = "TCP"

func (api *API) PortFwd() {
	n := node.GetNode(false)

	controllerEndpoint := getNodeControllerEndpoint(&topology.NodeReq{
		AccountID: n.AccountID,
		TenantID:  n.TenantID,
		NetID:     n.NetID,
		SubnetID:  n.SubnetID,
		NodeID:    n.NodeID,
	})

	nxnc, grpcConn := grpc.GetNetworkAPIClient(controllerEndpoint)
	defer grpcConn.Close()

	helpText := "Remote TCP port on target node"
	rp := input.GetInput("Remote TCP Port:", helpText, "", input.ValidPort)

	dstP, err := strconv.Atoi(rp)
	if err != nil {
		grpcConn.Close()
		status.Error(err, "Invalid destination port")
	}

	helpText = "Local TCP port"
	lp := input.GetInput("Local TCP Port:", helpText, "", input.ValidPort)

	srcP, err := strconv.Atoi(lp)
	if err != nil {
		grpcConn.Close()
		status.Error(err, "Invalid local port")
	}

	// rewrite mm.id for mmp authentication
	accountID, err := auth.GetAccountID()
	if err != nil {
		status.Error(fmt.Errorf("missing accountID"), "Unable to get account")
	}
	hostID := viper.GetString("host.id")
	mmID := fmt.Sprintf("%s:__cli:%s:%d:%d", accountID, hostID, os.Getegid(), time.Now().Unix())
	viper.Set("mm.id", mmID)

	srcID := viper.GetString("mm.id")

	nodeMMID, err := mm.GetID(&topology.NodeReq{
		AccountID: n.AccountID,
		TenantID:  n.TenantID,
		NetID:     n.NetID,
		SubnetID:  n.SubnetID,
		NodeID:    n.NodeID,
	})
	if err != nil {
		status.Error(err, "Unable to get node identifier")
	}
	dstID := nodeMMID.String()

	var wg sync.WaitGroup
	portFwdWaitc := make(chan struct{}, 2)

	wg.Add(1)
	go grpc.Control(nxnc, &wg, portFwdWaitc)

	authKey, err := auth.GetAuthKey()
	if err != nil {
		msg.Alert("Invalid or inexistent authorization key. Login to refresh your token.")
		grpcConn.Close()
		os.Exit(1)
	}

	output.SectionHeader("mmesh remote management")

	portfwd.NewPortFwd(authKey, uint32(srcP), uint32(dstP), proto, srcID, dstID, true)

	cli.Connected()

	output.CmdLog(fmt.Sprintf("Activating port forwarding to %s (%s/%s)", n.Cfg.NodeName, proto, rp))

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
		portfwd.ClosePortFwd(uint32(srcP), uint32(dstP), proto, srcID, dstID, &wg)

		wg.Wait()
		waitc <- struct{}{}
	}()
}
