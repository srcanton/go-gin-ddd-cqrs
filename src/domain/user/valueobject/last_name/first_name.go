package user_last_name

type LastName struct {
	value string
}

func FromValue(value string) (LastName, error) {
	return LastName{value: value}, nil
}

func (n LastName) GetValue() string {
	return n.value
}
