package acl

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

var aclsMap map[string]*iam.ACL = nil

func GetACL(edit bool) *iam.ACL {
	acll := ACLs()

	if len(acll) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var aclOptID string
	aclsOpts := make([]string, 0)
	acls := make(map[string]*iam.ACL)

	for _, acl := range acll {
		aclOptID = acl.ACLID
		aclsOpts = append(aclsOpts, aclOptID)
		acls[aclOptID] = acl
	}

	sort.Strings(aclsOpts)

	if edit {
		aclsOpts = append(aclsOpts, input.NewResource)
	}

	aclOptID = input.GetSelect("ACL:", "", aclsOpts, survey.Required)

	if aclOptID == input.NewResource {
		return nil
	}

	vars.ACLID = acls[aclOptID].ACLID

	return acls[aclOptID]
}

func ACLs() map[string]*iam.ACL {
	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetIAMAPIClient()
	defer grpcConn.Close()

	lr := &iam.ListACLsRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	acls := make(map[string]*iam.ACL)

	for {
		acll, err := nxc.ListACLs(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list ACLs")
		}

		for _, acl := range acll.ACLs {
			acls[acl.ACLID] = acl
		}

		if len(acll.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = acll.Meta.NextPageToken
		} else {
			break
		}
	}

	return acls
}

func validACLID(val interface{}) error {
	if err := input.ValidID(val); err != nil {
		return err
	}

	aclID := val.(string)

	if aclsMap == nil {
		aclsMap = ACLs()
	}

	if _, ok := aclsMap[aclID]; ok {
		return fmt.Errorf("ACL %s already exist", aclID)
	}

	return nil
}
