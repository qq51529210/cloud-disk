package db

type User struct {
	Id       string
	Name     string
	Password string
	State    int8
}

func GetUserByName(name string) (*User, error) {
	return nil, nil
}

func GetUserByPhone(name string) (*User, error) {
	return nil, nil
}
