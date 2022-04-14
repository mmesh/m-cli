package operation

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	workflow_pb "mmesh.dev/m-api-go/grpc/resources/ops/workflow"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-cli/pkg/client/ops/workflow"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-lib/pkg/mmid"
	"mmesh.dev/m-lib/pkg/utils"
)

func GetOperation() *workflow_pb.Operation {
	var opID string

	var opOpts []string
	opList := make(map[string]*workflow_pb.Operation)

	for _, o := range operations() {
		ref := strings.Split(o.OperationID, "@")[0]
		name := strings.Split(ref, ":")[0]
		tm, _ := utils.ParseTimestamp(strings.Split(ref, ":")[1])
		mmID := mmid.MMNodeID(o.NodeID)
		nodeID := mmID.NodeID()
		opID = fmt.Sprintf("Action: %s | Timestamp: %s\n  Target: %s\n", name, tm.String(), nodeID)
		opOpts = append(opOpts, opID)
		opList[opID] = o
	}

	sort.Strings(opOpts)

	opID = input.GetSelect("Action:", "", opOpts, survey.Required)

	return opList[opID]
}

func operations() []*workflow_pb.Operation {
	wf := workflow.GetWorkflow(false)

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	lr := &workflow_pb.ListOperationsRequest{
		Meta:     &resource.ListRequest{},
		Workflow: wf,
	}

	var operations []*workflow_pb.Operation

	for {
		ol, err := nxc.ListOperations(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list operations")
		}

		for _, o := range ol.Operations {
			operations = append(operations, o)
		}

		if len(ol.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = ol.Meta.NextPageToken
		} else {
			break
		}
	}

	return operations
}
