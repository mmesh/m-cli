package policy

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-cli/pkg/client/vrf"
)

func (api *API) Export() {
	v := vrf.GetVRF(false)

	fmt.Println()
	fmt.Println(networkPolicyExport(v))
	fmt.Println()
}

func networkPolicyExport(v *network.VRF) string {
	header := `# mmesh networkPolicy
apiVersion: v1
kind: NetworkPolicy

tenant: ` + v.TenantID + `
network: ` + v.NetID + `
subnet: ` + v.VRFID + `

enabled: true

networkPolicy:
  defaultPolicy: ` + v.NetworkPolicy.DefaultPolicy + `
  networkFilters:`

	var nfilters string
	for _, nf := range v.NetworkPolicy.NetworkFilters {
		f := `
  - index: ` + fmt.Sprintf("%d", nf.Index) + `
    description: ` + nf.Description + `
    srcIPNet: ` + nf.SrcIPNet + `
    dstIPNet: ` + nf.DstIPNet + `
    proto: ` + nf.Proto + `
    dstPort: ` + fmt.Sprintf("%d", nf.DstPort) + `
    policy: ` + nf.Policy

		nfilters += f
	}

	return header + nfilters
}
