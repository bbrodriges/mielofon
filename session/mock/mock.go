package mock

import (
	"github.com/bbrodriges/mielofon/v2/session"
)

var _ session.StoreSeeker = (*Store)(nil)

type Store struct {
	Values     map[string]interface{}
	MockGet    func(st Store, id string) (interface{}, error)
	MockSet    func(st Store, id string, sess interface{}) error
	MockDelete func(st Store, id string) error
	MockCount  func(st Store) int
}

func (s Store) Get(id string) (interface{}, error) {
	return s.MockGet(s, id)
}

func (s Store) Set(id string, sess interface{}) error {
	return s.MockSet(s, id, sess)
}

func (s Store) Delete(id string) error {
	return s.MockDelete(s, id)
}

func (s Store) Count() int {
	return s.MockCount(s)
}

func (s Store) VisitAll(f session.VisitFunc) {
	for k, v := range s.Values {
		if !f(k, v) {
			return
		}
	}
}
