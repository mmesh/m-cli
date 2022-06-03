package account

func (api *API) Show() {
	Output().Stats(fetchAccountStats())
}

/*
func (api *API) Stats() {
	Output().Stats(fetchAccountStats())
}
*/
