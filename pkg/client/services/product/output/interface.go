package output

import "mmesh.dev/m-api-go/grpc/resources/services"

type Interface interface {
	List(products map[string]*services.Product)
	Show(p *services.Product)

	ProductPrice(p *services.Product) string
	ProductSLA(sla services.ProductSLA) string
	ProductLocations(locations []string) string
	ProductLangs(langs []string) string
	ProviderRating(rating uint32) string
}
type API struct{}
