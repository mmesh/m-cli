package acl

import (
	"context"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) Set() {
	a := account.GetAccount()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	acl := GetACL(true)
	if acl != nil { // editing existing resource
		output.Choice("Edit RBAC ACL")
	} else { // <new> resource
		output.Choice("New RBAC ACL")

		acl = &iam.ACL{
			AccountID: a.AccountID,
			Users:     make(map[string]bool),
		}

		acl.ACLID = input.GetInput("ACL ID:", "", "", validACLID)
	}

	mmIDs, err := nxc.ListNodeMMIDs(context.TODO(), a)
	if err != nil {
		status.Error(err, "Unable to get node mmIDs")
	}
	acl.MMIDs = input.GetMultiSelect("Nodes:", "", mmIDs.MMIDs, acl.MMIDs, nil)

	s := output.Spinner()

	acl, err = nxc.SetACL(context.TODO(), acl)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to set ACL")
	}

	s.Stop()

	Output().Show(acl)
}
