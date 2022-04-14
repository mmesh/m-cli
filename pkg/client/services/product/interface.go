package product

import "mmesh.dev/m-cli/pkg/client/services/product/output"

type Interface interface {
	List(publicCatalog bool)
	Show(publicCatalog bool)
	Set(yamlFile string)
	Delete()
}
type API struct{}

func Resource() Interface {
	return &API{}
}

func Output() output.Interface {
	return &output.API{}
}
