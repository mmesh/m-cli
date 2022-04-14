package opportunity

import (
	"context"

	"github.com/AlecAivazis/survey/v2"
	"mmesh.dev/m-api-go/grpc/resources/services/platform/crm"
	"mmesh.dev/m-cli/pkg/grpc"
	"mmesh.dev/m-cli/pkg/input"
	"mmesh.dev/m-cli/pkg/output"
	"mmesh.dev/m-cli/pkg/status"
)

func (api *API) SetMilestone() {
	o := GetOpportunity()

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	omsr := &crm.OpportunityMilestoneSetRequest{
		AccountID:     o.AccountID,
		OpportunityID: o.OpportunityID,
		Milestone:     getMilestone(),
	}

	s := output.Spinner()

	o, err := nxc.SetOpportunityMilestone(context.TODO(), omsr)
	if err != nil {
		status.Error(err, "Unable to set opportunity milestone")
	}

	s.Stop()

	Output().Show(o)
}

func (api *API) SetOutcome() {
	o := GetOpportunity()

	nxc, grpcConn := grpc.GetServicesAPIClient(true)
	defer grpcConn.Close()

	outcome := getOutcome()

	var lostReason crm.OpportunityLostReason

	if outcome == crm.OpportunityOutcome_OUTCOME_LOST {
		lostReason = getLostReason()
	}

	s := output.Spinner()

	oosr := &crm.OpportunityOutcomeSetRequest{
		AccountID:     o.AccountID,
		OpportunityID: o.OpportunityID,
		Outcome:       outcome,
		LostReason:    lostReason,
	}

	o, err := nxc.SetOpportunityOutcome(context.TODO(), oosr)
	if err != nil {
		status.Error(err, "Unable to set opportunity outcome")
	}

	s.Stop()

	Output().Show(o)
}

func getMilestone() crm.OpportunityMilestone {
	milestones := map[string]crm.OpportunityMilestone{
		"NEW":         crm.OpportunityMilestone_NEW,
		"QUALIFIED":   crm.OpportunityMilestone_QUALIFIED,
		"MEETING":     crm.OpportunityMilestone_MEETING,
		"PROPOSAL":    crm.OpportunityMilestone_PROPOSAL,
		"NEGOTIATION": crm.OpportunityMilestone_NEGOTIATION,
		"CONTRACT":    crm.OpportunityMilestone_CONTRACT,
	}

	opts := make([]string, 0)

	for k := range milestones {
		opts = append(opts, k)
	}

	m := input.GetSelect("Opportunity Milestone:", "", opts, survey.Required)

	return milestones[m]
}

func getOutcome() crm.OpportunityOutcome {
	outcomes := map[string]crm.OpportunityOutcome{
		"Won":  crm.OpportunityOutcome_OUTCOME_WON,
		"Lost": crm.OpportunityOutcome_OUTCOME_LOST,
	}

	opts := make([]string, 0)

	for k := range outcomes {
		opts = append(opts, k)
	}

	o := input.GetSelect("Opportunity Outcome:", "", opts, survey.Required)

	return outcomes[o]
}

func getLostReason() crm.OpportunityLostReason {
	reasons := map[string]crm.OpportunityLostReason{
		"Unknown":            crm.OpportunityLostReason_LOST_UNKNOWN,
		"Abandoned":          crm.OpportunityLostReason_LOST_ABANDONED,
		"Became stale":       crm.OpportunityLostReason_LOST_BECAME_STALE,
		"Competition":        crm.OpportunityLostReason_LOST_COMPETITION,
		"No authority":       crm.OpportunityLostReason_LOST_NO_AUTHORITY,
		"Poor qualification": crm.OpportunityLostReason_LOST_POOR_QUALIFICATION,
	}

	opts := make([]string, 0)

	for k := range reasons {
		opts = append(opts, k)
	}

	r := input.GetSelect("Lost Reason:", "", opts, survey.Required)

	return reasons[r]
}
