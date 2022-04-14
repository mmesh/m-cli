package project

import (
	"context"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/ops/project"
	"mmesh.dev/m-api-go/grpc/resources/resource"
	"mmesh.dev/m-cli/pkg/client/account"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
	"mmesh.dev/m-cli/pkg/vars"
	"mmesh.dev/m-lib/pkg/utils/msg"
)

var projectsMap map[string]*project.Project = nil

func GetProject(edit bool) *project.Project {
	pl := projects()

	if len(pl) == 0 {
		msg.Info("No projects found")
		os.Exit(1)
	}

	var projectOptID string
	projectsOpts := make([]string, 0)
	projects := make(map[string]*project.Project)

	for _, p := range pl {
		projectOptID = p.ProjectID
		projectsOpts = append(projectsOpts, projectOptID)
		projects[projectOptID] = p
	}

	sort.Strings(projectsOpts)

	if edit {
		projectsOpts = append(projectsOpts, input.NewResource)
	}

	projectOptID = input.GetSelect("Project:", "", projectsOpts, survey.Required)

	if projectOptID == input.NewResource {
		return nil
	}

	vars.ProjectID = projects[projectOptID].ProjectID

	return projects[projectOptID]
}

func projects() map[string]*project.Project {
	a := account.GetAccount()

	s := output.Spinner()
	defer s.Stop()

	nxc, grpcConn := grpc.GetCoreAPIClient()
	defer grpcConn.Close()

	lr := &project.ListProjectsRequest{
		Meta:      &resource.ListRequest{},
		AccountID: a.AccountID,
	}

	projects := make(map[string]*project.Project)

	for {
		pl, err := nxc.ListProjects(context.TODO(), lr)
		if err != nil {
			status.Error(err, "Unable to list projects")
		}

		for _, p := range pl.Projects {
			projects[p.ProjectID] = p
		}

		if len(pl.Meta.NextPageToken) > 0 {
			lr.Meta.PageToken = pl.Meta.NextPageToken
		} else {
			break
		}
	}

	return projects
}

func validProjectID(val interface{}) error {
	if err := input.ValidID(val); err != nil {
		return err
	}

	projectID := val.(string)

	if projectsMap == nil {
		projectsMap = projects()
	}

	if _, ok := projectsMap[projectID]; ok {
		return fmt.Errorf("project %s already exist", projectID)
	}

	return nil
}
