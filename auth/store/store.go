package store

var (
	_User User
)

func Init(_type string, data map[string]interface{}) {
	switch _type {
	case "", "mongodb":

	}
}

func GetUser() User {
	return _User
}
