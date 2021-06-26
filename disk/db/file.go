package db

type File struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
}
