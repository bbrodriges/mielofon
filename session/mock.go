package session

var _ StoreSeeker = MockStore{}

type MockStore struct {
	Store      map[string]interface{}
	MockGet    func(st MockStore, id string) (interface{}, error)
	MockSet    func(st MockStore, id string, sess interface{}) error
	MockDelete func(st MockStore, id string) error
	MockCount  func(st MockStore) int
}

func (s MockStore) Get(id string) (interface{}, error) {
	return s.MockGet(s, id)
}

func (s MockStore) Set(id string, sess interface{}) error {
	return s.MockSet(s, id, sess)
}

func (s MockStore) Delete(id string) error {
	return s.MockDelete(s, id)
}

func (s MockStore) Count() int {
	return s.MockCount(s)
}

func (s MockStore) VisitAll(f VisitFunc) {
	for k, v := range s.Store {
		if !f(k, v) {
			return
		}
	}
}
