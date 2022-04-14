package sg

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/client/tenant"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Set() {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	sg := GetSecurityGroup(true)
	if sg != nil { // editing existing resource
		output.Choice("Edit RBAC Security Group")
	} else { // <new> resource
		output.Choice("New RBAC Security Group")

		sg = &iam.SecurityGroup{
			AccountID: a.AccountID,
			Users:     make(map[string]bool),
		}

		sg.SecurityGroupID = input.GetInput("Security Group ID:", "", "", validSecurityGroupID)
	}

	sg.Tenants = input.GetMultiSelect("Tenants:", "", getTenants(), sg.Tenants, nil)

	s := output.Spinner()

	sg, err := nxc.SetSecurityGroup(context.TODO(), sg)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set security-group")
	}

	s.Stop()

	Output().Show(sg)
}

func getTenants() []string {
	tl := tenant.Tenants()

	tenants := make([]string, 0)

	for _, t := range tl {
		tenants = append(tenants, t.TenantID)
	}

	return tenants
}
