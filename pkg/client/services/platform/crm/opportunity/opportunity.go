package opportunity

import (
	"context"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-api-go/grpc/resources/services/platform/crm"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func GetOpportunity() *crm.Opportunity {
	opportunities := opportunities()

	if len(opportunities) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	opportunitiesOpts := make([]string, 0)

	for opportunityID, _ := range opportunities {
		opportunitiesOpts = append(opportunitiesOpts, opportunityID)
	}

	opportunityID := input.GetSelect("Opportunity:", "", opportunitiesOpts, survey.Required)

	vars.OpportunityID = opportunityID

	return opportunities[opportunityID]
}

func opportunities() map[string]*crm.Opportunity {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	lr := &crm.ListOpportunitiesRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	opportunities := make(map[string]*crm.Opportunity)

	for {
		ol, err := nxc.ListOpportunities(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list CRM opportunities")
		}

		for _, o := range ol.Opportunities {
			opportunities[o.OpportunityID] = o
		}

		if len(ol.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = ol.Meta.NextPageToken
		} else {
			break
		}
	}

	return opportunities
}
