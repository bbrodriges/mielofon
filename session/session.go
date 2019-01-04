package session

import (
	"errors"
	"time"
)

var (
	ErrExpired = errors.New("session expired")
)

// TTLSession interface describes session object
// with lifetime
type TTLSession interface {
	SetTTL(time.Time) error
	Expired() bool
}

var _ TTLSession = new(ExpirableSession)

// ExpirableSession is a basic implamentation of
// TTLSession interface. It can be embedded in any
// other session struct
type ExpirableSession struct {
	ttl time.Time
}

func (s *ExpirableSession) SetTTL(ttl time.Time) error {
	if ttl.Before(time.Now()) {
		return ErrExpired
	}
	s.ttl = ttl
	return nil
}

func (s ExpirableSession) Expired() bool {
	return s.ttl.Before(time.Now())
}
