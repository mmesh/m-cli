package policy

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-api-go/grpc/resources/ops/object"
	"mmesh.dev/m-cli/pkg/client/node"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/ipnet"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

type ruleObjectType string

const (
	endpointObject  ruleObjectType = "Endpoint Object"
	ipNetCIDRObject ruleObjectType = "IPNet CIDR"
)

type ruleEndpoint struct {
	objectType ruleObjectType
	af         ipnet.AddressFamily
	addr       string
}

func setDefaultNetworkPolicy(v *network.VRF) *object.NetworkPolicyConfig {
	np := v.NetworkPolicy

	promptSelect := &survey.Select{Message: "Default Policy:", Options: vars.Policies}
	if err := survey.AskOne(promptSelect, &np.DefaultPolicy, survey.WithValidator(survey.Required), survey.WithIcons(input.SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}

	return &object.NetworkPolicyConfig{
		Tenant:        v.TenantID,
		Network:       v.NetID,
		Subnet:        v.VRFID,
		NetworkPolicy: np,
	}
}

func setNetworkPolicyRule(v *network.VRF) *object.NetworkPolicyConfig {
	np := v.NetworkPolicy

	i, nf := getNetworkFilter(np, true)
	if nf != nil && i != -1 { // editing existing resource
		output.Choice("Edit Object")
	} else { // <new> resource
		output.Choice("New Object")

		nf = &network.Filter{}
	}

	nf.Index = getPolicyIndex(nf.Index)

	prompt := &survey.Input{Message: "Description (optional):", Default: nf.Description}
	if err := survey.AskOne(prompt, &nf.Description, survey.WithIcons(input.SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}

	af := getAddressFamily()

	src := getRuleEndpoint(false, nf.SrcIPNet, af)
	nf.SrcIPNet = src.addr
	dst := getRuleEndpoint(true, nf.DstIPNet, af)
	nf.DstIPNet = dst.addr

	var protocols []string
	switch af {
	case ipnet.AddressFamilyIPv4:
		protocols = []string{"TCP", "UDP", "ICMPv4", "ANY"}
	case ipnet.AddressFamilyIPv6:
		protocols = []string{"TCP", "UDP", "ICMPv6", "ANY"}
	}
	promptSelect := &survey.Select{Message: "Protocol:", Options: protocols}
	if err := survey.AskOne(promptSelect, &nf.Proto, survey.WithValidator(survey.Required), survey.WithIcons(input.SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}

	if nf.Proto == "TCP" || nf.Proto == "UDP" {
		helpText := "Destination TCP or UDP port"
		defaultPort := fmt.Sprintf("%d", nf.DstPort)
		var dstP string
		prompt = &survey.Input{Message: "Destination Port:", Default: defaultPort, Help: helpText}
		if err := survey.AskOne(prompt, &dstP, survey.WithValidator(input.ValidPort), survey.WithIcons(input.SurveySetIcons)); err != nil {
			status.Error(err, "Unable to get response")
		}

		p, err := strconv.Atoi(dstP)
		if err != nil {
			status.Error(err, "Invalid port")
		}
		nf.DstPort = int32(p)
	}

	promptSelect = &survey.Select{Message: "Policy:", Options: vars.Policies}
	if err := survey.AskOne(promptSelect, &nf.Policy, survey.WithValidator(survey.Required), survey.WithIcons(input.SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}

	if np.NetworkFilters == nil {
		np.NetworkFilters = make([]*network.Filter, 0)
	}

	if i == -1 { // new element
		np.NetworkFilters = append(np.NetworkFilters, nf)
	} else {
		np.NetworkFilters[i] = nf
	}

	return &object.NetworkPolicyConfig{
		Tenant:        v.TenantID,
		Network:       v.NetID,
		Subnet:        v.VRFID,
		NetworkPolicy: np,
	}
}

func unsetNetworkPolicyRule(v *network.VRF) *object.NetworkPolicyConfig {
	np := v.NetworkPolicy

	i, _ := getNetworkFilter(np, false)

	if np.NetworkFilters == nil {
		np.NetworkFilters = make([]*network.Filter, 0)
	}

	// Remove the element at index i from np.NetworkFilters
	copy(np.NetworkFilters[i:], np.NetworkFilters[i+1:])             // Shift a[i+1:] left one index.
	np.NetworkFilters[len(np.NetworkFilters)-1] = nil                // Erase last element (write nil value).
	np.NetworkFilters = np.NetworkFilters[:len(np.NetworkFilters)-1] // Truncate slice.

	return &object.NetworkPolicyConfig{
		Tenant:        v.TenantID,
		Network:       v.NetID,
		Subnet:        v.VRFID,
		NetworkPolicy: np,
	}
}

func getNetworkFilter(np *network.Policy, edit bool) (int, *network.Filter) {
	var objects []string
	var objID string

	nfilters := make(map[string]int)

	for i, nf := range np.NetworkFilters {
		obj := fmt.Sprintf("%d: src %s | dst %s | %d/%s | %s", nf.Index, nf.SrcIPNet, nf.DstIPNet, nf.DstPort, nf.Proto, nf.Policy)
		objects = append(objects, obj)
		nfilters[obj] = i
	}

	sort.Strings(objects)
	if edit {
		objects = append(objects, input.NewResource)
	}

	if len(objects) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	promptSelect := &survey.Select{Message: "Network Policy:", Options: objects}
	if err := survey.AskOne(promptSelect, &objID, survey.WithValidator(survey.Required), survey.WithIcons(input.SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}

	if objID == input.NewResource {
		return -1, nil
	}

	i := nfilters[objID]

	return i, np.NetworkFilters[i]
}

func ipNetHelp() string {
	text := `CIDR notation and prefix lenght, like '192.168.1.14/32', '10.32.4.0/24',
'fd77:f:a45:2d2:1::/128' or '2001:db8::/32' as defined in RFC 4632 and RFC 4291`

	return text
}

func getPolicyIndex(defIdx uint32) uint32 {
	var index string

	prompt := &survey.Input{Message: "Policy Index:", Default: fmt.Sprintf("%d", defIdx)}
	if err := survey.AskOne(prompt, &index, survey.WithValidator(input.ValidUint), survey.WithIcons(input.SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}
	i, err := strconv.Atoi(index)
	if err != nil {
		status.Error(err, "Invalid index")
	}

	return uint32(i)
}

func getRuleEndpoint(dst bool, def string, af ipnet.AddressFamily) *ruleEndpoint {
	var ruleEndp *ruleEndpoint

	vars.NodeID = ""

	endpMsg := "Source"
	if dst {
		endpMsg = "Destination"
	}

	var endpObj string
	endpOpts := []string{string(endpointObject), string(ipNetCIDRObject)}
	endpSelect := &survey.Select{
		Message: "Select " + endpMsg + ":",
		Options: endpOpts,
	}
	if err := survey.AskOne(endpSelect, &endpObj, survey.WithValidator(survey.Required), survey.WithIcons(input.SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}

	if endpObj == string(endpointObject) {
		n := node.GetNode(false)
		e := node.GetEndpoint(n)
		ruleEndp = &ruleEndpoint{
			objectType: endpointObject,
			af:         af,
		}
		switch af {
		case ipnet.AddressFamilyIPv4:
			ruleEndp.addr = fmt.Sprintf("%s/32", e.IPv4)
		case ipnet.AddressFamilyIPv6:
			ruleEndp.addr = fmt.Sprintf("%s/128", e.IPv6)
		}
	} else { // endpObj == ipNetCIDRObject
		var ok bool
		var addr string
		prompt := &survey.Input{
			Message: endpMsg + " " + af.String() + " CIDR:",
			Default: def,
			Help:    ipNetHelp(),
		}
		for !ok {
			if err := survey.AskOne(prompt, &addr, survey.WithValidator(input.ValidIPNetCIDR), survey.WithIcons(input.SurveySetIcons)); err != nil {
				status.Error(err, "Unable to get response")
			}
			switch af {
			case ipnet.AddressFamilyIPv4:
				if strings.Contains(addr, ":") {
					msg.Error("Invalid IPv4 address")
					continue
				}
			case ipnet.AddressFamilyIPv6:
				if strings.Contains(addr, ".") {
					msg.Error("Invalid IPv6 address")
					continue
				}
			}
			ok = true
		}

		ruleEndp = &ruleEndpoint{
			objectType: ipNetCIDRObject,
			addr:       addr,
			af:         af,
		}
	}

	return ruleEndp
}

func getAddressFamily() ipnet.AddressFamily {
	var afOpt string

	afSelect := &survey.Select{
		Message: "Protocol:",
		Options: []string{
			ipnet.AddressFamilyIPv4.String(),
			ipnet.AddressFamilyIPv6.String(),
		},
	}
	if err := survey.AskOne(afSelect, &afOpt, survey.WithValidator(survey.Required), survey.WithIcons(input.SurveySetIcons)); err != nil {
		status.Error(err, "Unable to get response")
	}

	switch afOpt {
	case ipnet.AddressFamilyIPv4.String():
		return ipnet.AddressFamilyIPv4
	case ipnet.AddressFamilyIPv6.String():
		return ipnet.AddressFamilyIPv6
	}

	return ipnet.AddressFamilyUnspec
}
