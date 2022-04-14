package user

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-api-go/grpc/rpc"
	"mmesh.dev/m-cli/pkg/client/iam/acl"
	"mmesh.dev/m-cli/pkg/client/iam/role"
	"mmesh.dev/m-cli/pkg/client/iam/sg"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) SetPermissions() {
	u := GetUser(false)

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	supr := &iam.SetUserPermissionsRequest{
		AccountID: u.AccountID,
		Email:     u.Email,
		RBAC:      setUserRBAC(nxc, u),
	}

	s := output.Spinner()

	u, err := nxc.SetUserPermissions(context.TODO(), supr)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set user permissions")
	}

	s.Stop()

	Output().Show(u)
}

func setUserRBAC(nxc rpc.CoreAPIClient, u *iam.User) *iam.UserRBAC {
	rbac := u.RBAC

	rbac.SecurityGroups = input.GetMultiSelect("Security Groups:", "", getSecurityGroups(u.AccountID), rbac.SecurityGroups, nil)

	rbac.ACLs = input.GetMultiSelect("ACLs:", "", getACLs(u.AccountID), rbac.ACLs, nil)

	rbac.Roles = input.GetMultiSelect("Roles:", "", getRoles(u.AccountID), rbac.Roles, nil)

	return rbac
}

func getACLs(accountID string) []string {
	acll := acl.ACLs()

	acls := make([]string, 0)

	for _, acl := range acll {
		if acl.AccountID != accountID {
			continue
		}
		acls = append(acls, acl.ACLID)
	}

	return acls
}

func getRoles(accountID string) []string {
	rl := role.Roles()

	roles := make([]string, 0)

	for _, r := range rl {
		if r.AccountID != accountID {
			continue
		}
		roles = append(roles, r.RoleID)
	}

	return roles
}

func getSecurityGroups(accountID string) []string {
	sgl := sg.SecurityGroups()

	sgs := make([]string, 0)

	for _, sg := range sgl {
		if sg.AccountID != accountID {
			continue
		}
		sgs = append(sgs, sg.SecurityGroupID)
	}

	return sgs
}
