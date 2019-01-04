package session

import (
	"sync"
)

type MemoryStore struct {
	sync.Map
}

func NewMemoryStore(opts ...SeekerOption) (*MemoryStore, error) {
	s := new(MemoryStore)
	for _, opt := range opts {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (s *MemoryStore) Get(id string) (interface{}, error) {
	if sess, ok := s.Load(id); ok {
		return sess, nil
	}
	return nil, ErrNotFound
}

func (s *MemoryStore) Set(id string, sess interface{}) error {
	s.Store(id, sess)
	return nil
}

func (s *MemoryStore) Delete(id string) error {
	s.Map.Delete(id)
	return nil
}

func (s *MemoryStore) Count() int {
	var count int
	s.Range(func(k, v interface{}) bool {
		count++
		return true
	})
	return count
}

func (s *MemoryStore) VisitAll(f VisitFunc) {
	s.Range(func(id, sess interface{}) bool {
		return f(id.(string), sess)
	})
}
