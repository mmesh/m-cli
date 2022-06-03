package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(tenantID, netID, vrfID string, np *network.Policy) {
	output.SectionHeader("Security Policy Details")
	output.TitleT1("Security Policy")

	t := table.New()

	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(tenantID))
	t.AddRow(colors.Black("Network ID"), colors.DarkWhite(netID))
	t.AddRow(colors.Black("Subnet ID"), colors.DarkWhite(vrfID))

	t.Render()

	ShowNetworkPolicy(np)
}

func ShowNetworkPolicy(np *network.Policy) {
	t := table.New()

	switch np.DefaultPolicy {
	case ipnet.Policy_ACCEPT:
		t.AddRow(colors.Black("Default Policy"), output.StrEnabled(np.DefaultPolicy))
	case ipnet.Policy_DROP:
		t.AddRow(colors.Black("Default Policy"), output.StrDisabled(np.DefaultPolicy))
	}

	t.Render()
	fmt.Println()

	t = table.New()
	// t.Header(colors.Black("Index"), colors.Black("Source"), colors.Black("Destination"), colors.Black("Port/Proto"), colors.Black("Policy"))

	// t.AddRow(output.TableHeader("Index"), output.TableHeader("Source"), output.TableHeader("Destination"), output.TableHeader("Port/Proto"), output.TableHeader("Policy"))

	t.AddRow(colors.Black("Index"), colors.Black("Source"), colors.Black("Destination"), colors.Black("Port/Proto"), colors.Black("Policy"))
	t.AddRow(colors.Black("-----"), colors.Black("------"), colors.Black("-----------"), colors.Black("----------"), colors.Black("------"))

	for _, nf := range np.NetworkFilters {
		var pol string
		switch nf.Policy {
		case ipnet.Policy_ACCEPT:
			pol = output.StrEnabled(nf.Policy)
		case ipnet.Policy_DROP:
			pol = output.StrDisabled(nf.Policy)
		}

		var portProto string
		if nf.DstPort == 0 {
			portProto = nf.Proto
		} else {
			portProto = fmt.Sprintf("%d/%s", nf.DstPort, nf.Proto)
		}

		t.AddRow(
			fmt.Sprintf("%d", nf.Index),
			nf.SrcIPNet,
			nf.DstIPNet,
			portProto,
			pol,
		)
	}

	t.Render()
	fmt.Println()
}
