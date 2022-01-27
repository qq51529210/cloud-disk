package mongodb

type UserModel struct {
	ID       string
	Account  string
	Password string
	Phone    string
}

type User interface {
	Get(account string) (*UserModel, error)
}
