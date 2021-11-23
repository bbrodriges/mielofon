package session

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
)

// Store interface describes basic session store object
type Store interface {
	// Get returns session from store
	Get(id string) (interface{}, error)
	// Set saves session to store
	Set(id string, sess interface{}) error
	// Delete removes session from store
	Delete(id string) error
	// Count returns number of sessions in store
	Count() int
}

// StoreSeeker interface describes basic session store with
// additional ability to seek through all sessions
type StoreSeeker interface {
	Store
	// VisitAll visits all session records in store and applies VisitFunc to each one
	VisitAll(VisitFunc)
}

// VisitFunc is a function that will be applied to each session record in store
// via VisitAll function call. If VisitFunc returns false seek must be stopped
type VisitFunc func(id string, sess interface{}) bool
