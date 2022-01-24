package get_products

import aggregate "github.com/srcanton/go-gin-ddd-cqrs/src/domain/product/aggregate"

type GetProductQueryResponse struct {
	ID   string  `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

func NewGetProductsQueryResponse(products *[]aggregate.Product) *[]GetProductQueryResponse {
	var responseProducts []GetProductQueryResponse
	for _, product := range *products {
		responseProducts = append(responseProducts, GetProductQueryResponse{
			ID:   product.ID.GetValue(),
			Name: product.Name.GetValue(),
			Cost: product.Cost.GetValue(),
		})
	}
	return &responseProducts
}
