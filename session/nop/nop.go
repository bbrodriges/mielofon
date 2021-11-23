package nop

import (
	"github.com/bbrodriges/mielofon/v2/session"
)

var _ session.StoreSeeker = (*Store)(nil)

type Store struct{}

func (s Store) Get(_ string) (interface{}, error) {
	return nil, nil
}

func (s Store) Set(_ string, _ interface{}) error {
	return nil
}

func (s Store) Delete(_ string) error {
	return nil
}

func (s Store) Count() int {
	return 0
}

func (s Store) VisitAll(_ session.VisitFunc) {
	return
}
