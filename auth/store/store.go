package store

var (
	_UserStroe UserStroe
)

type Store interface {
	UserStroe() UserStroe
}

func SetUserStroe(st UserStroe) {
	_UserStroe = st
}

func GetUserStroe() UserStroe {
	return _UserStroe
}
