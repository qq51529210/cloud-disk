package mongodb

import "github.com/qq51529210/micro-services/auth/store"

func New(cfg map[string]interface{}) store.Store {
	st := new(Store)
	return st
}

type Store struct {
}

func (s *Store) UserStore() store.UserStore {
	return nil
}
