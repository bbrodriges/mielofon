package memory

import (
	"sync"

	"github.com/bbrodriges/mielofon/v2/session"
)

var _ session.StoreSeeker = (*Store)(nil)

// Store represents in-memory session store
type Store struct {
	sync.Map
}

// NewStore returns new in-memory session store instance
func NewStore() (*Store, error) {
	return new(Store), nil
}

func (s *Store) Get(id string) (interface{}, error) {
	if sess, ok := s.Load(id); ok {
		return sess, nil
	}
	return nil, session.ErrNotFound
}

func (s *Store) Set(id string, sess interface{}) error {
	s.Store(id, sess)
	return nil
}

func (s *Store) Delete(id string) error {
	s.Map.Delete(id)
	return nil
}

func (s *Store) Count() int {
	var count int
	s.Range(func(k, v interface{}) bool {
		count++
		return true
	})
	return count
}

func (s *Store) VisitAll(f session.VisitFunc) {
	s.Range(func(id, sess interface{}) bool {
		return f(id.(string), sess)
	})
}
