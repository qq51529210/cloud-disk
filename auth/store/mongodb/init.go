package mongodb

import "github.com/qq51529210/micro-services/auth/store"

func Init(cfg map[string]interface{}) store.Store {
	st := new(Store)
	return st
}
