package node

import (
	"context"
	"fmt"
	"os"

	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/client/vrf"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) AddNode() {
	v := vrf.GetVRF(false)
	if v == nil {
		msg.Alert("No subnet found.")
		msg.Alert("Please, configure at least one before creating networks.")
		os.Exit(1)
	}

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	/*
			fmt.Printf(`
		This command does NOT create a new node, just generates YAML
		config that you can use as a template for your convenience.

		The configuration comes with some important features disabled
		that you may want to enable (relay, kubernetes controller..).

		It is highly recommended to review the management section and
		customize it according to your security standards or criteria.

		If you feel you need some support, don't hesitate to contact us
		at %s ;-)

		`, colors.White("slack.mmesh.io"))
	*/

	vars.NodeID = input.GetInput("Node ID:", "", vars.NodeID, input.ValidID)

	niReq := &network.NodeInstanceRequest{
		AccountID: v.AccountID,
		TenantID:  v.TenantID,
		NetID:     v.NetID,
		VRFID:     v.VRFID,
		NodeID:    vars.NodeID,
	}

	s := output.Spinner()

	ni, err := nxc.CreateNodeInstallLinuxWebhook(context.TODO(), niReq)
	if err != nil {
		status.Error(err, "Unable to generate node config")
	}

	s.Stop()

	cmd1 := colors.DarkGreen(fmt.Sprintf("curl -s %s | sudo sh", ni.Config.WebhookURL))
	cmd2 := colors.DarkGreen(fmt.Sprintf("ssh <your_node> \"curl -s %s | sudo sh\"", ni.Config.WebhookURL))
	configFile := colors.DarkWhite("/etc/mmesh/mmesh-node.yml")

	fmt.Printf(`
The magic link generated is exclusively to connect the node %s to
the subnet %s. If you want to add more nodes, simply repeat this
process for them :-)

For security, these magic links expire in 24h and can be used once only,
but you can generate as many as you need.

Use the following commands to setup the node on the machine %s you
want to connect to your mmesh:

%s

or

%s

Once installed you can review (or modify) the configuration
at %s. It is recommended to review and customize
the management section according to your security standards or specific needs.

`,
		colors.DarkWhite(vars.NodeID),
		colors.DarkWhite(v.VRFID),
		colors.DarkWhite(vars.NodeID),
		cmd1,
		cmd2,
		configFile,
	)

	// fmt.Printf("\n# Generated config, use it to create your mmesh-agent.yml file\n")
}
