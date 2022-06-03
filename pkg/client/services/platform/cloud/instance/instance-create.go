package instance

import (
	"context"
	"fmt"

	"mmesh.dev/m-api-go/grpc/resources/iam"
	"mmesh.dev/m-api-go/grpc/resources/iam/auth"
	"mmesh.dev/m-api-go/grpc/resources/network"
	"mmesh.dev/m-api-go/grpc/resources/services"
	"mmesh.dev/m-api-go/grpc/resources/services/catalog/cloud/compute"
	"mmesh.dev/m-api-go/grpc/resources/services/platform/cloud"
	"mmesh.dev/m-cli/pkg/client/iam/user"
	"mmesh.dev/m-cli/pkg/client/provider"
	"mmesh.dev/m-cli/pkg/client/vrf"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

func (api *API) Create(loginReq *auth.LoginRequest, imgClass compute.ImageClass) {
	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	v := vrf.GetVRF(false)

	p := provider.GetProvider(services.ProviderType_CLOUD)

	img := provider.GetImage(p, imgClass)
	instanceType := provider.GetInstanceType(p, img)
	location := provider.GetLocation(p, img, instanceType)

	if img == nil || instanceType == nil || location == nil {
		status.Error(fmt.Errorf("VM instances not supported yet by %s integration", p.Name), "Unable to create resource")
	}

	vars.NodeID = input.GetInput("Node ID:", "", vars.NodeID, input.ValidID)

	sshKeys := make(map[string]*iam.SSHKey)

	k := user.Resource().GetSSHKey(loginReq)
	if k != nil {
		sshKeys[k.Name] = k
	}

	ciReq := &cloud.InstanceRequest{
		AccountID: v.AccountID,
		VRF:       v,
		NodeInstanceRequest: &network.NodeInstanceRequest{
			AccountID: v.AccountID,
			TenantID:  v.TenantID,
			NetID:     v.NetID,
			VRFID:     v.VRFID,
			NodeID:    vars.NodeID,
		},
		Spec: &cloud.InstanceSpec{
			Location:     location,
			InstanceType: instanceType,
			Image:        img,
			SSHKeys:      sshKeys,
		},
	}

	msg.Info("This action might take 1-5 minutes...")

	s := output.Spinner()

	ci, err := nxc.CreateCloudInstance(context.TODO(), ciReq)
	if err != nil {
		s.Stop()
		status.Error(err, "Unable to create cloud instance")
	}

	s.Stop()

	Output().Show(ci)

	if imgClass == compute.ImageClass_APP {
		fmt.Println(`Your cloud app is getting ready, but depending on the instance type
you chose, the process might take 5-20 minutes.
Please stick around and check with 'mmeshctl node list' when the node is ready :-)`)
	} else {
		fmt.Println(`Your cloud instance is getting ready, but depending on the instance type
you chose, the process might take 1-3 minutes.
Please stick around and check with 'mmeshctl node list' when the node is up :-)`)
	}
	fmt.Println()
}
