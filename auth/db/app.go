package db

type App struct {
	Id   string
	Name string
	Key  string
	Alg  string
}

func GetAppJwtInfo(id string) (*App, error) {
	return nil, nil
}
