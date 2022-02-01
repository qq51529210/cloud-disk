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
	AddUser(user *UserModel) (string, error)
	DeleteUser(_id string) (int64, error)
	UpdateUserPassword(model *UserModel) (int64, error)
	GetUser(account string) (*UserModel, error)
	GetUserList(query *PageQueryModel) (*PageDataModel, error)
}
