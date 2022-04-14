package instance

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-api-go/grpc/resources/services/catalog/cloud/compute"
	"mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func GetInstance(edit bool, imgClass compute.ImageClass) *cloud.Instance {
	il := instances(imgClass)

	if len(il) == 0 {
		msg.Info("No objects found")
		os.Exit(1)
	}

	var instanceOptID string
	instancesOpts := make([]string, 0)
	instances := make(map[string]*cloud.Instance)

	for _, i := range il {
		instanceOptID = fmt.Sprintf("[%s] %s", i.ProviderName, i.InstanceID)
		instancesOpts = append(instancesOpts, instanceOptID)
		instances[instanceOptID] = i
	}

	sort.Strings(instancesOpts)

	if edit {
		instancesOpts = append(instancesOpts, input.NewResource)
	}

	instanceOptID = input.GetSelect("Instance:", "", instancesOpts, survey.Required)

	if instanceOptID == input.NewResource {
		return nil
	}

	vars.CloudInstanceID = instances[instanceOptID].InstanceID

	return instances[instanceOptID]
}

func instances(imgClass compute.ImageClass) map[string]*cloud.Instance {
	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	lr := &cloud.ListInstancesRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	instances := make(map[string]*cloud.Instance)

	for {
		il, err := nxc.ListCloudInstances(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list cloud instances")
		}

		for _, i := range il.Instances {
			if i.Class == imgClass {
				instances[i.InstanceID] = i
			}
		}

		if len(il.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = il.Meta.NextPageToken
		} else {
			break
		}
	}

	return instances
}
