package thirdParty

import "mmesh.dev/m-api-go/grpc/resources/services/thirdParty"

type Interface interface {
	ClickUp(c *thirdParty.ClickUp) *thirdParty.ClickUp
	GitHub(gh *thirdParty.GitHub) *thirdParty.GitHub
	PagerDuty(pd *thirdParty.PagerDuty) *thirdParty.PagerDuty
	Slack(slck *thirdParty.Slack) *thirdParty.Slack
	// Crisp(crisp *thirdParty.Crisp) *thirdParty.Crisp

	DigitalOcean(dgo *thirdParty.DigitalOcean) *thirdParty.DigitalOcean
	GCP(dgo *thirdParty.GCP) *thirdParty.GCP
	Scaleway(scw *thirdParty.Scaleway) *thirdParty.Scaleway
	// AWS(aws *thirdParty.AWS) *thirdParty.AWS
}
type API struct{}

func Setup() Interface {
	return &API{}
}
