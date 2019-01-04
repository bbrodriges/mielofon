package session

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSetTTL(t *testing.T) {
	sess := new(ExpirableSession)
	err := sess.SetTTL(time.Now().Add(-1 * time.Second))
	assert.Equal(t, ErrExpired, err)

	sess = new(ExpirableSession)
	ttl := time.Now().Add(2 * time.Second)
	err = sess.SetTTL(ttl)
	assert.NoError(t, err)
	assert.Equal(t, ttl, sess.ttl)
}
