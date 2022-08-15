package account

import (
	"fmt"

	"mmesh.dev/m-cli/pkg/input"
)

func (api *API) Show() {
	a := fetchAccountStats()

	Output().Stats(a)

	if checkLimit(a) {
		return
	}

	if !input.GetConfirm("Upgrade subscription now?", true) {
		fmt.Println()
		return
	}

	api.BillingPortal(a)
}

/*
func (api *API) Stats() {
	Output().Stats(fetchAccountStats())
}
*/
