package output

import (
	"fmt"
	"os"
	"strings"

	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func agentInfo(n *network.Node) {
	na := n.Agent

	if na == nil {
		msg.Alert("No data found, agent is not registered yet")
		os.Exit(0)
	}

	if na.Metrics == nil {
		fmt.Printf("[%s]\n\n", colors.Black("not enough data yet"))
		os.Exit(0)
	}

	if na.Metrics.HostMetrics == nil {
		fmt.Printf("[%s]\n\n", colors.Black("not enough data yet"))
		os.Exit(0)
	}

	t := table.New()

	t.AddRow(colors.Black("Account ID"), colors.DarkWhite(n.AccountID))
	t.AddRow(colors.Black("Tenant ID"), colors.DarkWhite(n.TenantID))
	t.AddRow(colors.Black("Network ID"), colors.DarkWhite(n.NetID))
	t.AddRow(colors.Black("Subnet ID"), colors.DarkWhite(n.VRFID))
	t.AddRow(colors.Black("Node ID"), colors.DarkWhite(n.NodeID))
	// t.AddRow()
	// t.AddRow(colors.Black("Hostname"), colors.DarkWhite(na.AgentID))

	var status string
	if na.Healthy {
		status = output.StrOnline()
	} else {
		status = output.StrOffline()
	}
	t.AddRow(colors.Black("Status"), status)
	t.AddRow(colors.Black("OS"), colors.DarkWhite(strings.Title(na.Metrics.HostMetrics.OS)))

	if len(na.Metrics.HostMetrics.Uptime) > 0 {
		t.AddRow(colors.Black("Uptime"), colors.DarkWhite(na.Metrics.HostMetrics.Uptime))
	}

	var autoUpdate, maintSched string
	if na.Maintenance != nil {
		if na.Maintenance.AutoUpdate {
			autoUpdate = output.StrEnabled("auto-update")
		} else {
			autoUpdate = output.StrDisabled("auto-update")
		}
		if na.Maintenance.Schedule != nil {
			maintHour := fmt.Sprintf("%02d", na.Maintenance.Schedule.Hour)
			maintMin := fmt.Sprintf("%02d", na.Maintenance.Schedule.Minute)
			maintSched = colors.DarkGreen(fmt.Sprintf("%s:%s", maintHour, maintMin))
			maintSched = fmt.Sprintf("%s [%s]", colors.DarkWhite("Scheduled"), maintSched)
		} else {
			maintSched = colors.DarkWhite("Not scheduled")
		}
	}
	maint := fmt.Sprintf("%s %s", autoUpdate, maintSched)
	t.AddRow(colors.Black("Maintenance"), maint)

	var disableExec, disableTransfer, disablePortFwd, disableOps string
	if na.Management != nil {
		if na.Management.DisableExec {
			disableExec = output.StrDisabled("exec")
		} else {
			disableExec = output.StrEnabled("exec")
		}
		if na.Management.DisableTransfer {
			disableTransfer = output.StrDisabled("transfer")
		} else {
			disableTransfer = output.StrEnabled("transfer")
		}
		if na.Management.DisablePortForwarding {
			disablePortFwd = output.StrDisabled("portForward")
		} else {
			disablePortFwd = output.StrEnabled("portForward")
		}
		if na.Management.DisableOps {
			disableOps = output.StrDisabled("workflows")
		} else {
			disableOps = output.StrEnabled("workflows")
		}
	}
	opts := fmt.Sprintf("%s %s %s %s", disableExec, disableTransfer, disablePortFwd, disableOps)
	t.AddRow(colors.Black("Options"), opts)

	// t.AddRow()

	var extIP string
	if len(na.ExternalIPv4) > 0 {
		extIP = na.ExternalIPv4
	} else {
		extIP = "n/a"
	}
	t.AddRow(colors.Black("External IP"), colors.DarkWhite(extIP))
	t.AddRow(colors.Black("Port"), colors.DarkWhite(fmt.Sprintf("%d", na.Port)))

	var k8sGw string
	dnsPort := "n/a"
	if na.Options != nil {
		if na.Options.KubernetesGw {
			k8sGw = output.StrEnabled("k8sGw")
		}
		dnsPort = fmt.Sprintf("udp/%d", na.Options.DNSPort)
	}
	var relay string
	if na.IsRelay {
		relay = output.StrTier1()
	}
	prio := fmt.Sprintf("[%s-%s]", colors.DarkWhite("priority"), colors.Yellow(fmt.Sprintf("%d", na.Priority)))
	routing := fmt.Sprintf("%s %s %s", prio, relay, k8sGw)
	t.AddRow(colors.Black("DNS Port"), colors.DarkWhite(dnsPort))
	t.AddRow(colors.Black("Routing"), routing)

	// t.AddRow()

	hostResources := fmt.Sprintf("Load %s | RAM %s | Disk %s",
		colors.DarkWhite(fmt.Sprintf("%f", na.Metrics.HostMetrics.LoadAvg)),
		output.PercentBar(int(na.Metrics.HostMetrics.MemoryUsage)),
		output.PercentBar(int(na.Metrics.HostMetrics.DiskUsage)))
	t.AddRow(colors.Black("Resources"), hostResources)

	// t.AddRow(colors.Black("Memory"), output.PercentBar(int(na.Metrics.HostMetrics.MemoryUsage)))
	// t.AddRow(colors.Black("Load Average"), colors.DarkWhite(fmt.Sprintf("%f", na.Metrics.HostMetrics.LoadAvg)))
	// t.AddRow(colors.Black("Disk"), output.PercentBar(int(na.Metrics.HostMetrics.DiskUsage)))

	t.Render()
	fmt.Println()
}
