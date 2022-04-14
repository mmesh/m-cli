package alert

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/events"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func GetAlert() *events.Alert {
	al := alerts()

	if len(al) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var alertOptID string
	alertsOpts := make([]string, 0)
	alerts := make(map[string]*events.Alert)

	for _, a := range al {
		tm := output.Datetime(a.LastUpdated)
		nodeID := output.Fit(a.NodeID, 32)
		status := strings.ToUpper(a.Status)
		// component := output.Fit(a.Component, 32)
		component := strings.ToLower(a.Component)
		l1 := fmt.Sprintf("%s %s [%s]", tm, nodeID, status)
		// l2 := fmt.Sprintf("Component: %s", component)
		l2 := fmt.Sprintf("Component: %s | Group: %s | Class: %s", component, a.Group, a.Class)
		// l3 := fmt.Sprintf("Group: %s | Class: %s", a.Group, a.Class)
		// alertOptID = fmt.Sprintf("%s\n  %s\n  %s\n", l1, l2, l3)
		alertOptID = fmt.Sprintf("%s\n  %s\n", l1, l2)
		alertsOpts = append(alertsOpts, alertOptID)
		alerts[alertOptID] = a
	}

	sort.Strings(alertsOpts)

	alertOptID = input.GetSelect("Alert:", "", alertsOpts, survey.Required)

	vars.AlertID = alerts[alertOptID].AlertID

	return alerts[alertOptID]
}

func alerts() map[string]*events.Alert {
	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	lr := &events.ListAlertsRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	alerts := make(map[string]*events.Alert)

	for {
		al, err := nxc.ListAlerts(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list alerts")
		}

		for _, a := range al.Alerts {
			alerts[a.AlertID] = a
		}

		if len(al.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = al.Meta.NextPageToken
		} else {
			break
		}
	}

	return alerts
}
