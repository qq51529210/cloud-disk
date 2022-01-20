package store

type UserModel struct {
	ID       string
	Account  string
	Password string
	Phone    string
}

type UserStore interface {
	Get(account string) (*UserModel, error)
}
