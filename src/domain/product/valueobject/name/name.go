package product_name

type ProductName struct {
	value string
}

func FromValue(value string) (ProductName, error) {
	return ProductName{value: value}, nil
}

func (n ProductName) GetValue() string {
	return n.value
}
