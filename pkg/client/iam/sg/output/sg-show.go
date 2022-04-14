package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(sg *iam.SecurityGroup) {
	output.SectionHeader("IAM: Security Group Details")
	output.TitleT1("Security Group Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(sg.AccountID))
	t.AddRow(colors.Black("Security Group"), colors.DarkWhite(sg.SecurityGroupID))

	t.Render()
	fmt.Println()

	if len(sg.Tenants) > 0 {
		output.SubTitleT2("Tenants")

		for _, tenant := range sg.Tenants {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(tenant))
		}
		fmt.Println()
	}

	if len(sg.Users) > 0 {
		output.SubTitleT2("Users")

		for u := range sg.Users {
			fmt.Printf(" ■ %s\n", colors.DarkGreen(u))
		}
		fmt.Println()
	}
}
