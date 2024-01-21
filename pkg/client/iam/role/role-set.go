package role

import (
	"context"

	"mmesh.dev/m-api-go/grpc/common/empty"
	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Set() {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	role := GetRole(true)
	if role != nil { // editing existing resource
		output.Choice("Edit RBAC Role")
	} else { // <new> resource
		output.Choice("New RBAC Role")

		role = &iam.Role{
			AccountID: a.AccountID,
			Users:     make(map[string]bool),
		}

		role.RoleID = input.GetInput("Role ID:", "", "", validRoleID)
	}

	perms, err := nxc.ListIAMPermissions(context.TODO(), &empty.Request{})
	if err != nil {
		status.Error(err, "Unable to get IAM permissions")
	}

	var permList []string
	for _, p := range perms.Permissions {
		permList = append(permList, p)
	}

	role.Permissions = input.GetMultiSelect("Permissions:", "", permList, role.Permissions, nil)

	s := output.Spinner()

	role, err = nxc.SetRole(context.TODO(), role)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set role")
	}

	s.Stop()

	Output().Show(role)
}
