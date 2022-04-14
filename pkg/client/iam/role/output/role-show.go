package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(r *iam.Role) {
	output.SectionHeader("IAM: Role Details")
	output.TitleT1("Role Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(r.AccountID))
	t.AddRow(colors.Black("Role"), colors.DarkWhite(r.RoleID))

	t.Render()
	fmt.Println()

	if len(r.Permissions) > 0 {
		output.SubTitleT2("Permissions")

		for _, perm := range r.Permissions {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(perm))
		}
		fmt.Println()
	}

	if len(r.Users) > 0 {
		output.SubTitleT2("Users")

		for u := range r.Users {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(u))
		}
		fmt.Println()
	}
}
