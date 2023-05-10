package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-api-go/grpc/resources/topology"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(acl *iam.ACL, nsm *topology.NodeSummaryMap) {
	output.SectionHeader("IAM: ACL Details")
	output.TitleT1("ACL Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(acl.AccountID))
	t.AddRow(colors.Black("ACL"), colors.DarkWhite(acl.ACLID))

	t.Render()

	fmt.Println()

	if len(acl.NodeIDs) > 0 {
		output.SubTitleT2("Allowed Nodes")

		for _, nodeID := range acl.NodeIDs {
			if n, ok := nsm.Nodes[nodeID]; ok {
				nodeStr := fmt.Sprintf("[%s] %s: %s", n.TenantName, n.SubnetID, n.NodeName)
				fmt.Printf(" ■ %s\n", colors.DarkGreen(nodeStr))
			}
		}
		fmt.Println()
	}

	if len(acl.Users) > 0 {
		output.SubTitleT2("Users")

		for u := range acl.Users {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(u))
		}
		fmt.Println()
	}
}
