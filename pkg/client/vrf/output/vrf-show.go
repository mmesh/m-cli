package output

import (
	"fmt"
	"strconv"

	"mmesh.dev/m-api-go/grpc/resources/network"
	np_output "mmesh.dev/m-cli/pkg/client/policy/output"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) Show(v *network.VRF) {
	output.SectionHeader("Subnet Details")
	output.TitleT1("Subnet Information")

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(v.AccountID))
	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(v.TenantID))
	t.AddRow(colors.Black("Network ID"), colors.DarkWhite(v.NetID))
	t.AddRow(colors.Black("Subnet ID"), colors.DarkWhite(v.VRFID))
	t.AddRow(colors.Black("Description"), colors.DarkWhite(v.Description))

	t.Render()
	fmt.Println()

	output.SubTitleT2("Credentials")

	fmt.Printf(colors.Black("Subnet Authorization Secret %s\n\n"), colors.DarkWhite("******************************"))

	fmt.Print(colors.Black("-----SUBNET AUTHORIZATION KEY-----\n"))
	fmt.Printf("%s\n", colors.DarkWhite(v.AuthKey.Key))
	fmt.Print(colors.Black("-----SUBNET AUTHORIZATION KEY-----\n"))
	fmt.Println()

	if v.IPAM != nil {
		output.SubTitleT2("IP Address Management (IPAM)")

		t := table.New()

		t.AddRow(colors.Black("Subnet CIDR"), colors.DarkWhite(v.IPAM.SubnetCIDR))
		t.AddRow(colors.Black("IPv4 Addresses Available"), colors.DarkWhite(strconv.Itoa(int(v.IPAM.TotalAvailable))))
		t.AddRow(colors.Black("IPv4 Addresses Leased"), colors.DarkWhite(strconv.Itoa(int(v.IPAM.TotalLeased))))

		t.Render()
		fmt.Println()

		if v.IPAM.TotalLeased > 0 {
			t := table.New()
			// t.Header(colors.Black("IPv4 ADDRESS"), colors.Black("IPv6 ADDRESS"), colors.Black("ENDPOINT"))
			t.Header(colors.Black("IPv4 Address"), colors.Black("IPv6 Address"), colors.Black("Endpoint"))
			// t.Header(output.TableHeader("IPv4 Address"), output.TableHeader("IPv6 Address"), output.TableHeader("Endpoint"))

			for ipv4, leaseEndpointMap := range v.IPAM.Leased {
				ipv6, err := ipnet.GetIPv6(ipv4)
				if err != nil {
					ipv6 = "-"
				}

				l := 0
				for endpoint := range leaseEndpointMap.Endpoints {
					if l == 0 {
						t.AddRow(colors.DarkWhite(ipv4), ipv6, colors.DarkWhite(output.Fit(endpoint, 32)))
					} else {
						t.AddRow("", "", colors.DarkWhite(endpoint))
					}
					l++
				}
			}

			t.Render()
			fmt.Println()
		}
	}

	if v.NetworkPolicy != nil {
		output.SubTitleT2("Network Policy")
		np_output.ShowNetworkPolicy(v.NetworkPolicy)
	}

	// if v.RelayService != nil {
	// 	output.SubTitleT2("Managed Relay Service")

	// 	t = uitable.New()
	// 	t.MaxColWidth = 60
	// 	t.Wrap = false

	// 	if v.RelayService.Node != nil {
	// 		t.AddRow("Node ID:", colors.DarkWhite(v.RelayService.Node.NodeID))

	// 		l := v.RelayService.Location
	// 		location := fmt.Sprintf("%s/%s", l.Region, l.Zone)
	// 		t.AddRow("Location:", colors.DarkWhite(location))
	// 	}

	// 	var enabled, suspended string
	// 	if v.RelayService.Suspended {
	// 		suspended = "[" + colors.Red("SUSPENDED") + "]"
	// 	}
	// 	if v.RelayService.Enabled {
	// 		enabled = "[" + colors.Green("ENABLED") + "]"
	// 	} else {
	// 		enabled = "[" + colors.Red("DISABLED") + "]"
	// 	}
	// 	t.AddRow("Status:", enabled+" "+suspended)

	// 	fmt.Println(t)

	// 	fmt.Println()
	// }
}
