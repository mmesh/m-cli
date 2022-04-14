package output

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gosuri/uitable"
	"mmesh.dev/m-api-go/grpc/network/mmnp/routing"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(rt *routing.RoutingTable) {
	output.SectionHeader("mmesh routing engine")
	output.TitleT1("Routing Table")

	t := uitable.New()
	t.MaxColWidth = 60
	t.Wrap = false

	// connRoute := fmt.Sprintf(" %s%s%s", colors.Black("["), colors.Magenta("C"), colors.Black("]"))
	intRoute := fmt.Sprintf(" %s%s%s", colors.Black("["), colors.Magenta("I"), colors.Black("]"))
	extRoute := fmt.Sprintf(" %s%s%s", colors.Black("["), colors.Magenta("E"), colors.Black("]"))
	legend := "        Legend:"

	// CONNECTED not included
	numRoutes := len(rt.RE) - 2
	if numRoutes < 0 {
		numRoutes = 0
	}
	t.AddRow("Scope:", colors.DarkWhite(rt.Scope.String()), legend, intRoute+" Internal route")
	t.AddRow("Routes:", colors.DarkWhite(fmt.Sprintf("%d", numRoutes)), "", extRoute+" External route")
	// t.AddRow("Relays:", colors.DarkWhite(fmt.Sprintf("%d", len(rt.Relays))))

	fmt.Println(t)
	fmt.Println()

	t = uitable.New()
	t.MaxColWidth = 48
	t.Wrap = false

	for route, re := range rt.RE {
		// if re.Type == routing.RouteType_CONNECTED {
		// 	continue
		// }

		if strings.HasPrefix(route, ipnet.MMesh64Prefix) {
			continue
		}

		nGw := colors.DarkCyan(strconv.Itoa(len(re.Gw)))
		gateways := fmt.Sprintf("%s %s", colors.Black("mmesh gateways:"), colors.DarkCyan(nGw))

		var ha string
		if len(re.Gw) > 1 {
			ha = fmt.Sprintf("%s %s", colors.Black("High Availability:"), colors.Green("yes"))
		} else {
			ha = fmt.Sprintf("%s %s", colors.Black("High Availability:"), colors.DarkRed("no"))
		}

		switch re.Type {
		// case routing.RouteType_CONNECTED:
		// 	t.AddRow(colors.White(route) + connRoute)
		case routing.RouteType_DYNAMIC:
			ipv6, err := ipnet.GetIPv6(strings.Split(route, "/")[0])
			if err == nil && len(ipv6) > 0 {
				addrv6 := fmt.Sprintf("%s/128", ipv6)
				if _, ok := rt.RE[addrv6]; ok {
					t.AddRow(colors.DarkWhite(route) + intRoute)
					t.AddRow(colors.DarkWhite(addrv6), ha+" | "+gateways)
				}
			} else {
				t.AddRow(colors.DarkWhite(route)+extRoute, ha+" | "+gateways)
			}
		case routing.RouteType_STATIC:
			t.AddRow(colors.DarkWhite(route)+extRoute, ha+" | "+gateways)
		}

		for prio, nHop := range re.Gw {
			var conns int
			var lastSeen string
			if nHop.LinkStatus != nil {
				conns = int(nHop.LinkStatus.Connections)

				if nHop.LinkStatus.LastSeen == 0 {
					lastSeen = colors.Black("n/a")
				} else {
					tm := time.Unix(nHop.LinkStatus.LastSeen, 0)
					lastSeen = colors.DarkWhite(tm.String())
				}
			}

			p := fmt.Sprintf("prio-%d", prio)
			nhPrio := fmt.Sprintf("%s%s%s", colors.Yellow("["), colors.DarkYellow(p), colors.Yellow("]"))
			// nhPrio := fmt.Sprintf("  [prio-%s]", colors.Red(strconv.Itoa(int(prio))))

			nhConns := fmt.Sprintf("Connections: %s", colors.Yellow(strconv.Itoa(conns)))

			var nhHealthy string
			if nHop.Healthy {
				nhHealthy = fmt.Sprintf("Status: %s", output.StrOnline())
			} else {
				nhHealthy = fmt.Sprintf("Status: %s", output.StrOffline())
			}

			var relay string
			if nHop.Agent.IsRelay {
				relay = output.StrTier1()
			}

			t.AddRow(nhPrio, colors.DarkYellow("-> ")+"GwID: "+colors.DarkWhite(nHop.Agent.AgentID)+" "+relay)
			// t.AddRow("", "   HostID: "+colors.DarkWhite(nHop.Agent.NxHostID))
			t.AddRow("", "   Last Seen: "+lastSeen)
			t.AddRow("", "   "+nhHealthy+" | "+nhConns)
			// t.AddRow()
			// for _, maddr := range nHop.Agent.MAddrs {
			// 	t.AddRow("", colors.DarkYellow(" -> ")+colors.Black(maddr), colors.DarkWhite(" "))
			// }
		}
		t.AddRow()
	}
	fmt.Println(t)

	fmt.Printf("%s %s\n", colors.Black("Last Updated"), colors.DarkWhite(time.Unix(rt.LastUpdated, 0).String()))

	fmt.Println()

	// output.SubTitleT2("Subnet Relays")

	// rl := uitable.New()
	// rl.MaxColWidth = 72
	// rl.Wrap = false

	// rs := fmt.Sprintf(" %s%s%s", colors.DarkYellow("["), colors.Yellow("R"), colors.DarkYellow("]"))
	// for _, r := range rt.Relays {
	// 	rl.AddRow(rs, colors.Black(r.MAddr))
	// }
	// fmt.Println(rl)
}
