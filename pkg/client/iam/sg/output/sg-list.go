package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(securityGroups map[string]*iam.SecurityGroup) {
	output.SectionHeader("IAM: Security Groups")
	output.TitleT1("Security Group List")

	t := table.New()
	t.Header(colors.Black("ACCOUNT ID"), colors.Black("SECURITY GROUP"))

	for _, sg := range securityGroups {
		t.AddRow(sg.AccountID, colors.DarkWhite(sg.SecurityGroupID))
	}

	t.Render()
	fmt.Println()
}
