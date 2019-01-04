package session

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var _ TTLSession = new(testSession)

type testSession struct {
	ExpirableSession

	id     string
	userId string
	ttl    time.Time
}

func TestWithVacuum(t *testing.T) {
	opt := WithVacuum(500 * time.Millisecond)
	ms, err := NewMemoryStore(opt)

	assert.NoError(t, err)

	sess := &testSession{
		id:     "2eac4854-fce721f3-b845abba-20d60",
		userId: "0A5BDE6A080BD274CC5E4825ED5F384FD5FF94629DC025C013EA793DDCD62EE2",
		ExpirableSession: ExpirableSession{
			ttl: time.Now().Add(1 * time.Second),
		},
	}
	assert.False(t, sess.Expired())

	err = ms.Set("2eac4854-fce721f3-b845abba-20d60", sess)
	assert.NoError(t, err)

	assert.Equal(t, 1, ms.Count())
	time.Sleep(1500 * time.Millisecond)
	assert.Equal(t, 0, ms.Count())
}
