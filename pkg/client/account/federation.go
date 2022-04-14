package account

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/controller"
	location_pb "mmesh.dev/m-api-go/grpc/resources/location"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func selectFederation() *controller.FederationSelected {
	ll := selectLocation()

	var lOptions []string
	locations := make(map[string]*location_pb.Location)

	for _, l := range ll {
		lID := fmt.Sprintf("%s / %s", l.Region, l.Zone)
		lOptions = append(lOptions, lID)
		locations[lID] = l
	}

	sort.Strings(lOptions)

	if len(lOptions) == 0 {
		msg.Alert("No location found")
		os.Exit(1)
	}

	lID := vars.LocationID

	if len(lID) == 0 {
		lID = input.GetSelect("Location:", "", lOptions, survey.Required)
	}

	l, ok := locations[lID]
	if !ok {
		msg.Alert("Invalid selection")
		os.Exit(1)
	}

	vars.LocationID = lID

	nxc, grpcConn := grpc.GetManagerProviderAPIClient(false)
	defer grpcConn.Close()

	s := output.Spinner()

	f, err := nxc.SelectFederation(context.TODO(), l)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to get federation")
	}

	s.Stop()

	return f
}
