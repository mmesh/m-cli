package vs

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

var vsMap map[string]*topology.VS = nil
var selectedVS *topology.VS = nil

func getVS(edit bool) *topology.VS {
	if selectedVS != nil {
		return selectedVS
	}

	vsl := vss()

	if len(vsl) == 0 {
		msg.Info("No virtual server found")
		os.Exit(1)
	}

	var vsOptID string
	vssOpts := make([]string, 0)
	vss := make(map[string]*topology.VS)

	for _, vs := range vsl {
		vsOptID = fmt.Sprintf("[%s] %s", vs.Name, vs.Description)
		vssOpts = append(vssOpts, vsOptID)
		vss[vsOptID] = vs
	}

	sort.Strings(vssOpts)

	if edit {
		vssOpts = append(vssOpts, input.NewResource)
	}

	vsOptID = input.GetSelect("VS:", "", vssOpts, survey.Required)

	if vsOptID == input.NewResource {
		return nil
	}

	vars.VSID = vss[vsOptID].VSID
	selectedVS = vss[vsOptID]

	return vss[vsOptID]
}

func vss() map[string]*topology.VS {
	if vsMap != nil {
		return vsMap
	}

	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetTopologyAPIClient()
	defer grpcConn.Close()

	lr := &topology.ListVSsRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	vss := make(map[string]*topology.VS)

	for {
		vsl, err := nxc.ListVSs(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list virtual servers")
		}

		for _, vs := range vsl.VSs {
			vss[vs.VSID] = vs
		}

		if len(vsl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = vsl.Meta.NextPageToken
		} else {
			break
		}
	}

	vsMap = vss

	return vss
}
