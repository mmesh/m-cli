package cluster

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func GetKubernetesCluster(edit bool) *cloud.KubernetesCluster {
	kcl := k8sClusters()

	if len(kcl) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var k8sClusterOptID string
	k8sClustersOpts := make([]string, 0)
	k8sClusters := make(map[string]*cloud.KubernetesCluster)

	for _, kc := range kcl {
		k8sClusterOptID = fmt.Sprintf("[%s] %s", kc.ProviderID, kc.KubernetesClusterID)
		k8sClustersOpts = append(k8sClustersOpts, k8sClusterOptID)
		k8sClusters[k8sClusterOptID] = kc
	}

	sort.Strings(k8sClustersOpts)

	if edit {
		k8sClustersOpts = append(k8sClustersOpts, input.NewResource)
	}

	k8sClusterOptID = input.GetSelect("Cluster:", "", k8sClustersOpts, survey.Required)

	if k8sClusterOptID == input.NewResource {
		return nil
	}

	vars.CloudKubernetesClusterID = k8sClusters[k8sClusterOptID].KubernetesClusterID

	return k8sClusters[k8sClusterOptID]
}

func k8sClusters() map[string]*cloud.KubernetesCluster {
	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	lr := &cloud.ListKubernetesClustersRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	k8sClusters := make(map[string]*cloud.KubernetesCluster)

	for {
		kcl, err := nxc.ListCloudKubernetesClusters(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list cloud kubernetes clusters")
		}

		for _, kc := range kcl.KubernetesClusters {
			k8sClusters[kc.KubernetesClusterID] = kc
		}

		if len(kcl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = kcl.Meta.NextPageToken
		} else {
			break
		}
	}

	return k8sClusters
}
