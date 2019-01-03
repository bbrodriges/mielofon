package session

import (
	"time"
)

type SeekerOption func(StoreSeeker) error

// WithVacuum starts dedicated goroutine and on each interval
// iterates over store, searching sessions expired by TTL and
// deletes them
func WithVacuum(interval time.Duration) SeekerOption {
	return func(s StoreSeeker) error {
		go func() {
			ticker := time.NewTicker(interval)
			defer ticker.Stop()

			for range ticker.C {
				s.VisitAll(func(id string, sess interface{}) bool {
					if ts, ok := sess.(TTLSession); ok && ts.Expired() {
						_ = s.Delete(id)
					}
					return true
				})
			}
		}()
		return nil
	}
}
