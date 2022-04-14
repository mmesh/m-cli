package output

import (
	"fmt"
	"strings"

	"mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(instances map[string]*cloud.Instance) {
	output.SectionHeader("Cloud: Instances")
	output.TitleT1("Instance List")

	t := table.New()
	t.Header(colors.Black("CLOUD PROVIDER"), colors.Black("INSTANCE ID"))

	for _, i := range instances {
		providerID := output.StrNormal(strings.Title(i.ProviderID))
		t.AddRow(providerID, colors.DarkWhite(output.Fit(i.InstanceID, 32)))
	}

	t.Render()
	fmt.Println()
}
