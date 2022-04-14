package vars

import "mmesh.dev/m-lib/pkg/ipnet"

var YAMLFile string

var PricingPlanID string
var LocationID string
var AccountID, TenantID, NetID, VRFID, NodeID string
var ProjectID, WorkflowID, OperationID string
var UserEmail, SecurityGroupID, RoleID, ACLID string
var AlertID string
var InvoiceID, BilledItemID string
var ProviderID, ProductID string
var CloudInstanceID, CloudKubernetesClusterID string
var IssueID string
var OpportunityID string

// var ContractID string

var Policies = []string{
	ipnet.Policy_ACCEPT,
	ipnet.Policy_DROP,
}
