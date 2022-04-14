package rt

import (
	"context"

	"mmesh.dev/m-api-go/grpc/network/mmnp/routing"
	"mmesh.dev/m-cli/pkg/client/vrf"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) List() {
	v := vrf.GetVRF(false)

	s := output.Spinner()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	rt, err := nxc.ListRoutes(context.TODO(), v)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to list routes")
	}

	if rt.Scope == routing.Scope_VRF {
		for r, re := range rt.RE {
			if re.VRFID != v.VRFID {
				delete(rt.RE, r)
			}
		}
	}

	s.Stop()

	// output.Show(rt)
	Output().List(rt)
}
