package provider

import (
	"context"
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

func GetProvider(ptype services.ProviderType) *services.Provider {
	providers := ListProviders(ptype)

	if len(providers) == 0 {
		switch ptype {
		case services.ProviderType_PROFESSIONAL_SERVICES:
			msg.Info("No service provider found")
		default:
			msg.Info("No provider found")
		}
		os.Exit(1)
	}

	providerOpts := make([]string, 0)

	for providerID, _ := range providers {
		providerOpts = append(providerOpts, providerID)
	}

	providerID := input.GetSelect("Provider:", "", providerOpts, survey.Required)

	vars.ProviderID = providerID

	return providers[providerID]
}

func ListProviders(ptype services.ProviderType) map[string]*services.Provider {
	serviceID := viper.GetString("serviceID")

	s := output.Spinner()

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	lr := &services.ListProvidersRequest{
		Meta:         &resource.ListRequest{},
		ServiceID:    serviceID,
		ProviderType: ptype,
	}

	providers := make(map[string]*services.Provider)

	for {
		pl, err := nxc.ListProviders(context.TODO(), lr)
		if err != nil {
			s.Stop()
			status.Error(err, "Unable to list providers")
		}

		for _, p := range pl.Providers {
			if p.Type == ptype || ptype == services.ProviderType_UNSPECIFIED {
				providers[p.ProviderID] = p
			}
		}

		if len(pl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = pl.Meta.NextPageToken
		} else {
			break
		}
	}

	s.Stop()

	return providers
}
