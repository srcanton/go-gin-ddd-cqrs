package product_cost

type ProductCost struct {
	value float64
}

func FromValue(value float64) (ProductCost, error) {
	return ProductCost{value: value}, nil
}

func (n ProductCost) GetValue() float64 {
	return n.value
}
