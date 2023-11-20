package network

import (
	"context"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/location"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func getConnectivityZone() *location.Location {
	ll := listConnectivityLocations()

	if len(ll) == 0 {
		msg.Info("Connectivity zone not found")
		os.Exit(1)
	}

	var locationOptID string
	locationsOpts := make([]string, 0)
	locations := make(map[string]*location.Location)

	for _, l := range ll {
		locationOptID = l.LocationID

		if l.ConnectivityZone != nil {
			if l.ConnectivityZone.Active {
				locationsOpts = append(locationsOpts, locationOptID)
				locations[locationOptID] = l
				continue
			}
		}
	}

	sort.Strings(locationsOpts)

	locationOptID = input.GetSelect("Connectivity Zone:", "", locationsOpts, survey.Required)

	return locations[locationOptID]
}

func listConnectivityLocations() []*location.Location {
	nxc, grpcConn := grpc.GetManagerAPIClient(true)
	defer grpcConn.Close()

	lr := &location.ListLocationsRequest{
		Meta: &resource.ListRequest{},
	}

	var locations []*location.Location
	var ll *location.Locations = nil
	var err error

	for ll == nil || len(lr.Meta.PageToken) > 0 {
		ll, err = nxc.ListConnectivityLocations(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list connectivity-zone locations")
		}

		locations = append(locations, ll.Locations...)

		lr.Meta.PageToken = ll.Meta.NextPageToken
	}

	return locations
}
