package session

var _ StoreSeeker = NopStore{}

type NopStore struct{}

type NopSession struct{}

func (s NopStore) Get(_ string) (interface{}, error) {
	return NopSession{}, nil
}

func (s NopStore) Set(_ string, _ interface{}) error {
	return nil
}

func (s NopStore) Delete(_ string) error {
	return nil
}

func (s NopStore) Count() int {
	return 0
}

func (s NopStore) VisitAll(_ VisitFunc) {
	return
}