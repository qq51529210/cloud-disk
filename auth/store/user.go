package store

type UserModel struct {
	ID       string
	Account  string
	Password string
	Phone    string
}

type UserStroe interface {
	Get(account string) (*UserModel, error)
}
