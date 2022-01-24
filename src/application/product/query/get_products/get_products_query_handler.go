package get_products

func (bus *Bus) handleFindProductsQuery(
	query *GetProductsQuery,
) (*[]GetProductQueryResponse, error) {
	products, err := bus.service.GetProducts()
	if err != nil {
		return nil, err
	}
	return NewGetProductsQueryResponse(products), nil
}
