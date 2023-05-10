package command

import (
	"context"
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func getNodeControllerEndpoint(nr *topology.NodeReq) string {
	s := output.Spinner()

	nxc, grpcConn := grpc.GetControllerAPIClient()
	defer grpcConn.Close()

	c, err := nxc.GetNodeController(context.TODO(), nr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get nodeController")
	}

	s.Stop()

	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
