package session

var _ StoreSeeker = MockStore{}

type MockStore struct {
	Store      map[string]interface{}
	MockGet    func(id string) (interface{}, error)
	MockSet    func(id string, sess interface{}) error
	MockDelete func(id string) error
	MockCount  func() int
}

func (s MockStore) Get(id string) (interface{}, error) {
	return s.MockGet(id)
}

func (s MockStore) Set(id string, sess interface{}) error {
	return s.MockSet(id, sess)
}

func (s MockStore) Delete(id string) error {
	return s.MockDelete(id)
}

func (s MockStore) Count() int {
	return s.MockCount()
}

func (s MockStore) VisitAll(f VisitFunc) {
	for k, v := range s.Store {
		if !f(k, v) {
			return
		}
	}
}
