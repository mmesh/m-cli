package itsm

import (
	"context"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"mmesh.dev/m-api-go/grpc/resources/services/platform/crm"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/client/services/product"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) ServiceRequest() {
	a := account.GetAccount()

	userEmail := viper.GetString("user.email")
	if len(userEmail) == 0 {
		userEmail = input.GetInput("User Email:", "", "", input.ValidEmail)
	}

	userNickname := strings.Split(userEmail, "@")[0]

	company := a.AccountID // TODO: add companyName to acccount struct

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	serviceCategory := product.SelectServiceCategory()

	p := product.RequestProductService(serviceCategory)

	reqHelp := "Enter your requirements or custom details for the service provider"
	customerReqs := input.GetMultiline("Requirements/Details:", reqHelp, "", survey.Required)

	svcReq := &crm.ServiceRequest{
		AccountID:         p.ProviderID,
		CustomerAccountID: a.AccountID,
		CustomerEmail:     userEmail,
		CustomerNickname:  userNickname,
		CustomerCompany:   company,
		CustomerText:      customerReqs,
		Spec: &crm.ServiceSpec{
			Product: p,
		},
	}

	s := output.Spinner()

	i, err := nxc.RequestService(context.TODO(), svcReq)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to create service request")
	}

	s.Stop()

	Output().Show(i)
}
