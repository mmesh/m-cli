package services

import (
	"mmesh.dev/m-cli/pkg/client/services/product"
)

type Interface interface {
	Product() product.Interface
}
type API struct{}

func (a *API) Product() product.Interface {
	return &product.API{}
}
