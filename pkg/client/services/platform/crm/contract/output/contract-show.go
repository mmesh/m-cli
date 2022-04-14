package output

import (
	"fmt"
	"time"

	"mmesh.dev/m-api-go/grpc/resources/services/platform/crm"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(c *crm.Contract) {
	output.SectionHeader("Services: Contract Details")
	output.TitleT1("Contract Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(c.AccountID))
	t.AddRow(colors.Black("Contract ID"), colors.DarkWhite(c.ContractID))
	t.AddRow(colors.Black("Provider ID"), colors.DarkWhite(c.Spec.Product.ProviderID))
	t.AddRow(colors.Black("Product"), colors.DarkWhite(c.Spec.Product.Name))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(c.Spec.Product.Description))

	tm := time.Unix(c.CreationDate, 0).String()
	t.AddRow(colors.Black("Creation Date"), colors.DarkWhite(tm))
	tm = time.Unix(c.LastModified, 0).String()
	t.AddRow(colors.Black("Last Modified"), colors.DarkWhite(tm))

	t.Render()
	fmt.Println()
}
