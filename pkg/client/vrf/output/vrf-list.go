package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(vrfs map[string]*network.VRF) {
	output.SectionHeader("Subnets")
	output.TitleT1("Subnet List")

	t := table.New()
	t.Header(colors.Black("SUBNET ID"), colors.Black("SUBNET CIDR"), colors.Black("DESCRIPTION"))

	for _, v := range vrfs {
		if v.IPAM != nil {
			t.AddRow(v.VRFID, colors.DarkWhite(v.IPAM.SubnetCIDR), output.Fit(v.Description, 32))
		}
	}

	t.Render()
	fmt.Println()
}
