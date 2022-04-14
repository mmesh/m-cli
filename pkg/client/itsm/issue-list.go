package itsm

func (api *issueAPI) List() {
	Output().List(issues())
}
