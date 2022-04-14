package itsm

import "mmesh.dev/m-cli/pkg/client/itsm/output"

type Interface interface {
	Issue() IssueInterface

	Chat()

	// AssistanceInfo()
	// AssistanceSupport()

	// Feedback()

	// Incident()

	// Problem()

	// Change()

	// RequestUncategorized()
	// RequestAdvisoryService()
	// RequestSpecialProject()
	// RequestSpecialTask()
	// RequestSpecificFunctionality()
	// RequestExternalService()

	ServiceRequest()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
