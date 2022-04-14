package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(acls map[string]*iam.ACL) {
	output.SectionHeader("IAM: ACLs")
	output.TitleT1("ACL List")

	t := table.New()
	t.Header(colors.Black("ACCOUNT ID"), colors.Black("ACL"))

	for _, a := range acls {
		t.AddRow(a.AccountID, colors.DarkWhite(a.ACLID))
	}

	t.Render()
	fmt.Println()
}
