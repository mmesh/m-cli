package pricingplan

import (
	"context"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-api-go/grpc/resources/services"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func GetPricingPlan() *services.PricingPlan {
	ppl := ListPricingPlans(false)

	var ppOpt string
	var pricingPlanOpts []string
	pricingPlans := make(map[string]*services.PricingPlan)
	for _, pp := range ppl {
		ppOpt = fmt.Sprintf("%s", pp.Description)
		pricingPlanOpts = append(pricingPlanOpts, ppOpt)
		pricingPlans[ppOpt] = pp
	}

	if len(pricingPlanOpts) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	ppOpt = input.GetSelect("Plan:", "", pricingPlanOpts, survey.Required)

	vars.PricingPlanID = pricingPlans[ppOpt].PricingPlanID

	return pricingPlans[ppOpt]
}

func ListPricingPlans(reqAuth bool) []*services.PricingPlan {
	serviceID := viper.GetString("serviceID")

	s := output.Spinner()

	nxc, grpcConn := grpc.GetServicesAPIClient(reqAuth)
	defer grpcConn.Close()

	lr := &services.ListPricingPlansRequest{
		Meta:      &resource.ListRequest{},
		ServiceID: serviceID,
	}

	var pricingPlans []*services.PricingPlan

	for {
		pl, err := nxc.ListPricingPlans(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list pricingPlans")
		}

		pricingPlans = append(pricingPlans, pl.PricingPlans...)

		if len(pl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = pl.Meta.NextPageToken
		} else {
			break
		}
	}

	s.Stop()

	return pricingPlans
}
