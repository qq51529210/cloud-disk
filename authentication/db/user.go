package db

import "errors"

var (
	ErrUserExists = errors.New("username exists")
)

type User struct {
	Id       string
	Name     string
	Password string
	State    int8
}

func (u *User) StateString() string {
	switch u.State {
	case 0:
		return "enable"
	case 1:
		return "disable"
	default:
		return "invalid"
	}
}

func GetUserByName(name string) (*User, error) {

	return nil, nil
}

func GetUserByPhone(name string) (*User, error) {
	return nil, nil
}

func CreateUser(user *User) error {
	return nil
}
