package user_password

type Password struct {
	value string
}

func FromValue(value string) (Password, error) {
	return Password{value: value}, nil
}

func (n Password) GetValue() string {
	return n.value
}
