package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(users map[string]*iam.User) {
	output.SectionHeader("IAM: Users")
	output.TitleT1("User List")

	t := table.New()
	t.Header(colors.Black("ACCOUNT ID"), colors.Black("USER ID"))

	for _, u := range users {
		t.AddRow(u.AccountID, colors.DarkWhite(u.UserID))
	}

	t.Render()
	fmt.Println()
}
