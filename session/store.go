package session

import (
	"errors"
	"time"
)

var (
	ErrNotFound      = errors.New("session not found")
	ErrAlreadyExists = errors.New("session already exits")
)

// Store interface describes basic session store object
type Store interface {
	Get(id string) (interface{}, error)
	Set(id string, sess interface{}) error
	Delete(id string) error
	Count() int
}

// Store interface describes basic session store with
// additional ability to seek through all sessions
type StoreSeeker interface {
	Store
	VisitAll(VisitFunc)
}

type VisitFunc func(id string, sess interface{}) bool

var (
	ErrExpired = errors.New("session expired")
)

// TTLSession interface describes session object
// with lifetime
type TTLSession interface {
	SetTTL(time.Time) error
	Expired() bool
}
