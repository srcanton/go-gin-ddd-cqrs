package user_first_name

type FirstName struct {
	value string
}

func FromValue(value string) (FirstName, error) {
	return FirstName{value: value}, nil
}

func (n FirstName) GetValue() string {
	return n.value
}
