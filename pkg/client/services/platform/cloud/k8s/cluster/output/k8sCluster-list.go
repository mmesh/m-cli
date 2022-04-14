package output

import (
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/output/table"
	"mmesh.dev/m-lib/pkg/utils/colors"
)

func (api *API) List(k8sClusters map[string]*cloud.KubernetesCluster) {
	output.SectionHeader("Cloud: Kubernetes Clusters")
	output.TitleT1("Kubernetes Cluster List")

	t := table.New()
	t.Header(colors.Black("CLUSTER ID"), colors.Black("REGION"), colors.Black("VERSION"), colors.Black("STATUS"))

	for _, kc := range k8sClusters {
		kcID := colors.DarkWhite(kc.KubernetesClusterID)
		t.AddRow(kcID, kc.Region, kc.Version, k8sStatus(kc.Status))
	}

	t.Render()
	fmt.Println()
}
