package output

import (
	"fmt"
	"sort"

	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(nodes map[string]*network.Node) {
	output.SectionHeader("Nodes")
	output.TitleT1("Node List")

	t := table.New()
	t.Header(colors.Black("NODE ID / FQDN"), colors.Black("IPv4"), colors.Black("FLAGS / ENDPOINT"))

	t.SetRowLine("-")

	nodesSort := make([]string, 0)
	for _, n := range nodes {
		nodesSort = append(nodesSort, n.NodeID)
	}
	sort.Strings(nodesSort)

	for _, nodeID := range nodesSort {
		n := nodes[nodeID]

		var c1, c2, c3 string
		var status string

		nodeID := colors.Black(output.Fit(n.NodeID, 32))

		if n.Agent != nil {
			// hostTitle := fmt.Sprintf(" %s %s:", colors.DarkWhite("└─"), colors.Black("Node Instance"))
			p := fmt.Sprintf("prio-%d", n.Agent.Priority)
			prio := fmt.Sprintf("%s%s%s", colors.Yellow("["), colors.DarkYellow(p), colors.Yellow("]"))

			var nodeType string
			if n.Agent.IsMRS {
				nodeType = output.StrMRS()
			} else {
				if n.Agent.IsRelay {
					nodeType = output.StrTier1()
				}
			}

			if n.Agent.Healthy {
				status = colors.Green("█")
			} else {
				status = colors.DarkRed("█")
			}

			c1 = fmt.Sprintf("%s %s", status, nodeID)
			c2 = ""
			c3 = fmt.Sprintf("%s %s", prio, nodeType)
		}

		if n.Endpoints != nil {
			for _, e := range n.Endpoints {
				fqdn := colors.DarkWhite(output.Fit(e.DNSName+".mmesh.local", 32))
				endpoint := output.Fit(e.EndpointID, 20)
				ipv4 := colors.DarkWhite(e.IPv4)

				c1 = fmt.Sprintf("%s\n%s %s", c1, status, fqdn)
				c2 = fmt.Sprintf("%s\n%s", c2, ipv4)
				c3 = fmt.Sprintf("%s\n%s", c3, endpoint)
			}
		}
		t.AddRow(c1, c2, c3)
	}

	t.Render()
	fmt.Println()
}

/*
func (api *API) List2(nodes []*network.Node) {
	output.SectionHeader("Nodes")
	output.TitleT1("Node List")

	t := table.New()
	t.Header(colors.Black("NODE ID / FQDN"), colors.Black("IPv4"), colors.Black("FLAGS / ENDPOINT"))

	for _, n := range nodes {
		nodeID := output.Fit(n.NodeID, 32)

		if n.Agent != nil {
			// hostTitle := fmt.Sprintf(" %s %s:", colors.DarkWhite("└─"), colors.Black("Node Instance"))
			p := fmt.Sprintf("priority-%d", n.Agent.Priority)
			prio := fmt.Sprintf("%s%s%s", colors.Yellow("["), colors.DarkYellow(p), colors.Yellow("]"))

			var nodeType string
			if n.Agent.IsMRS {
				nodeType = output.StrMRS()
			} else {
				if n.Agent.IsRelay {
					nodeType = output.StrTier1()
				}
			}

			var status string
			if n.Agent.Healthy {
				status = output.StrEnabled("■")
			} else {
				status = output.StrDisabled("■")
			}

			t.AddRow(status+" "+nodeID, "", prio+" "+nodeType)
		}

		if n.Endpoints != nil {
			for _, e := range n.Endpoints {
				fqdn := colors.DarkWhite(output.Fit(e.DNSName+".mmesh.local", 32))
				endpoint := output.Fit(e.EndpointID, 20)
				t.AddRow("    "+fqdn, colors.DarkWhite(e.IPv4), endpoint)
			}
		}
		t.AddRow("", "", "")
	}

	t.Render()
	fmt.Println()
}
*/

/*
func (api *API) List2(nodes []*network.Node) {
	output.SectionHeader("Nodes")
	output.TitleT1("Node List")

	t := uitable.New()
	t.MaxColWidth = 24
	t.Wrap = false

	t.AddRow("NODE ID / ENDPOINT", "FLAGS / IPv4", "STATUS / FQDN")
	sep := strings.Repeat("─", 22)
	t.AddRow(sep, sep, sep)

	for _, n := range nodes {
		if n.Agent != nil {
			//hostTitle := fmt.Sprintf(" %s %s:", colors.DarkWhite("└─"), colors.Black("Node Instance"))
			prio := fmt.Sprintf("[priority %s]", colors.Yellow(strconv.Itoa(int(n.Agent.Priority))))

			var relay string
			if n.Agent.IsMRS {
				relay = output.StrMRS()
			} else {
				if n.Agent.IsRelay {
					relay = output.StrTier1()
				}
			}

			var status string
			if n.Agent.Healthy {
				//status = fmt.Sprintf("       [%s]", colors.Green("ONLINE"))
				status = fmt.Sprintf("       %s", output.StrOnline())
			} else {
				//status = fmt.Sprintf("       [%s]", colors.Red("OFFLINE"))
				status = fmt.Sprintf("       %s", output.StrOffline())
			}

			t.AddRow(n.NodeID, prio+" "+relay, status)
		}

		if n.Endpoints != nil {
			for _, e := range n.Endpoints {
				fqdn := e.DNSName + ".mmesh.local"
				t.AddRow(" "+colors.DarkWhite(e.EndpointID), colors.DarkWhite(e.IPv4), fqdn)
			}
		}
		t.AddRow()
	}
	fmt.Println(t)

	fmt.Println()
}
*/
