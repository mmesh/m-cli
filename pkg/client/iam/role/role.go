package role

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

var rolesMap map[string]*iam.Role = nil

func GetRole(edit bool) *iam.Role {
	rl := Roles()

	if len(rl) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var roleOptID string
	rolesOpts := make([]string, 0)
	roles := make(map[string]*iam.Role)

	for _, r := range rl {
		roleOptID = r.RoleID
		rolesOpts = append(rolesOpts, roleOptID)
		roles[roleOptID] = r
	}

	sort.Strings(rolesOpts)

	if edit {
		rolesOpts = append(rolesOpts, input.NewResource)
	}

	roleOptID = input.GetSelect("Role:", "", rolesOpts, survey.Required)

	if roleOptID == input.NewResource {
		return nil
	}

	vars.RoleID = roles[roleOptID].RoleID

	return roles[roleOptID]
}

func Roles() map[string]*iam.Role {
	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	lr := &iam.ListRolesRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	roles := make(map[string]*iam.Role)

	for {
		rl, err := nxc.ListRoles(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list roles")
		}

		for _, r := range rl.Roles {
			roles[r.RoleID] = r
		}

		if len(rl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = rl.Meta.NextPageToken
		} else {
			break
		}
	}

	return roles
}

func validRoleID(val interface{}) error {
	if err := input.ValidID(val); err != nil {
		return err
	}

	roleID := val.(string)

	if rolesMap == nil {
		rolesMap = Roles()
	}

	if _, ok := rolesMap[roleID]; ok {
		return fmt.Errorf("role %s already exist", roleID)
	}

	return nil
}
