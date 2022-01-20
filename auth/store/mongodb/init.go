package mongodb

import "github.com/qq51529210/micro-services/auth/store"

func New(cfg map[string]interface{}) store.Store {
	st := new(_Store)
	return st
}

type _Store struct {
}
