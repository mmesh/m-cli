package command

import (
	"context"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/account"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-api-go/grpc/resources/tenant"
	"mmesh.dev/m-cli/pkg/auth"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-lib/pkg/mmid"
	"mmesh.dev/m-lib/pkg/mmp"
	"mmesh.dev/m-lib/pkg/mmp/cli"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) Transfer(args []string) {
	nxnc, grpcConn := grpc.GetNetworkAPIClient()
	defer grpcConn.Close()

	var wg sync.WaitGroup
	transferWaitc := make(chan struct{}, 2)

	wg.Add(1)
	go grpc.Control(nxnc, &wg, transferWaitc)

	runTransfer(args)

	// transferWaitc <- struct{}{}

	wg.Wait()

}

func runTransfer(args []string) {
	var srcID, dstID string
	var srcFilePath, dstFilePath string
	var err error

	mmID := viper.GetString("mm.id")

	src := args[0]
	dst := args[1]

	s := strings.Split(src, ":")
	d := strings.Split(dst, ":")

	if strings.ToLower(s[0]) == "localhost" && strings.ToLower(d[0]) == "localhost" {
		msg.Error("Both src and dst cannot be local paths.")
		os.Exit(1)
	}

	// src
	if strings.ToLower(s[0]) == "localhost" && len(s) == 2 {
		srcID = mmID
		srcFilePath = s[1]
	} else {
		srcID, srcFilePath, err = checkTransferArg(src)
		if err != nil {
			msg.Errorf("Invalid source %s", src)
			os.Exit(1)
		}
	}

	// dst
	if strings.ToLower(d[0]) == "localhost" && len(d) == 2 {
		dstID = mmID
		dstFilePath = d[1]
	} else {
		dstID, dstFilePath, err = checkTransferArg(dst)
		if err != nil {
			msg.Errorf("Invalid destination %s", dst)
			os.Exit(1)
		}
	}

	if len(dstFilePath) == 0 {
		dstFilePath = srcFilePath
	}

	if len(srcID) > 0 && len(srcFilePath) > 0 && len(dstID) > 0 && len(dstFilePath) > 0 {
		authKey, err := auth.GetAuthKey()
		if err != nil {
			msg.Error("Invalid or inexistent authorization key. Login to refresh your token.")
			os.Exit(1)
		}

		output.SectionHeader("mmesh remote management")

		cli.Connected()

		output.CmdLog("Transferring files via mmesh network..")
		fmt.Println()

		mmp.NewTransfer(authKey, srcFilePath, dstFilePath, srcID, dstID, true)
	} else {
		msg.Error("Invalid arguments")
		os.Exit(1)
	}
}

func checkTransferArg(arg string) (string, string, error) {
	var filePath string

	s := strings.Split(arg, ":")

	if len(s) < 5 || len(s) > 6 {
		msg.Errorf("Invalid argument")
		return "", "", fmt.Errorf("Invalid nodeURI")
	}

	nxc1, grpcConn1 := grpc.GetControllerProviderAPIClient(true)
	defer grpcConn1.Close()

	accountID := s[0]
	tenantID := s[1]
	netID := s[2]
	vrfID := s[3]
	nodeID := s[4]

	if len(s) == 6 {
		filePath = s[5]
	}

	// check account
	a := &account.Account{
		AccountID: accountID,
	}

	if _, err := nxc1.GetAccount(context.TODO(), a); err != nil {
		msg.Errorf("Invalid account %s", accountID)
		return "", "", err
	}

	nxc2, grpcConn2 := grpc.GetCoreAPIClient()
	defer grpcConn2.Close()

	// check tenant
	t := &tenant.Tenant{
		AccountID: accountID,
		TenantID:  tenantID,
	}

	if _, err := nxc2.GetTenant(context.TODO(), t); err != nil {
		msg.Errorf("Invalid tenant %s", tenantID)
		return "", "", err
	}

	// check network
	netw := &network.Network{
		AccountID: accountID,
		TenantID:  tenantID,
		NetID:     netID,
	}

	if _, err := nxc2.GetNetwork(context.TODO(), netw); err != nil {
		msg.Errorf("Invalid network %s", netID)
		return "", "", err
	}

	// check vrf / subnet
	v := &network.VRF{
		AccountID: accountID,
		TenantID:  tenantID,
		NetID:     netID,
		VRFID:     vrfID,
	}

	if _, err := nxc2.GetVRF(context.TODO(), v); err != nil {
		msg.Errorf("Invalid subnet %s", vrfID)
		return "", "", err
	}

	// check node
	n := &network.Node{
		AccountID: accountID,
		TenantID:  tenantID,
		NetID:     netID,
		VRFID:     vrfID,
		NodeID:    nodeID,
	}

	if _, err := nxc2.GetNode(context.TODO(), n); err != nil {
		msg.Errorf("Invalid node %s", nodeID)
		return "", "", err
	}

	nodeMMID := mmid.GetMMIDFromNode(n).String()

	return nodeMMID, filePath, nil
}
