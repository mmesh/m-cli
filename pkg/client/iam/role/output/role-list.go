package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(roles map[string]*iam.Role) {
	output.SectionHeader("IAM: Roles")
	output.TitleT1("Role List")

	t := table.New()
	t.Header(colors.Black("ACCOUNT ID"), colors.Black("ROLE"))

	for _, r := range roles {
		t.AddRow(r.AccountID, colors.DarkWhite(r.RoleID))
	}

	t.Render()
	fmt.Println()
}
