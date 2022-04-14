package services

import (
	"mmesh.dev/m-cli/pkg/client/services/platform"
	"mmesh.dev/m-cli/pkg/client/services/product"
)

type Interface interface {
	Product() product.Interface
	Platform() platform.Interface
}
type API struct{}

func (a *API) Product() product.Interface {
	return &product.API{}
}

func (a *API) Platform() platform.Interface {
	return &platform.API{}
}
