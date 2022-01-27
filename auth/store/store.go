package store

var (
	_Store Store
)

func SetStore(store Store) {
	_Store = store
}

func GetStore() Store {
	return _Store
}

type Store interface {
	GetUser(account string) (*UserModel, error)
}
