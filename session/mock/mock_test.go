package mock

import (
	"errors"
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreGet(t *testing.T) {
	called := false
	ms := Store{
		MockGet: func(_ Store, _ string) (interface{}, error) {
			called = true
			return "shimba", nil
		},
	}
	sess, err := ms.Get("ololo")
	assert.Nil(t, err)
	assert.Equal(t, "shimba", sess)
	assert.True(t, called)

	called = false
	ms = Store{
		MockGet: func(_ Store, _ string) (interface{}, error) {
			called = true
			return nil, errors.New("unknown error")
		},
	}
	sess, err = ms.Get("ololo")
	assert.Nil(t, sess)
	assert.Error(t, err)
	assert.True(t, called)

	called = false
	ms = Store{
		Values: map[string]interface{}{"ololo": "shimba"},
		MockGet: func(s Store, id string) (interface{}, error) {
			called = true
			if sess, ok := s.Values[id]; ok {
				return sess, nil
			}
			return nil, errors.New("not found")
		},
	}
	sess, err = ms.Get("ololo")
	assert.Equal(t, "shimba", sess)
	assert.NoError(t, err)
	assert.True(t, called)
}

func TestStoreSet(t *testing.T) {
	called := false
	ms := Store{
		MockSet: func(_ Store, _ string, _ interface{}) error {
			called = true
			return nil
		},
	}
	err := ms.Set("ololo", "shimba")
	assert.Nil(t, err)
	assert.True(t, called)

	called = false
	ms = Store{
		MockSet: func(_ Store, _ string, _ interface{}) error {
			called = true
			return errors.New("unknown error")
		},
	}
	err = ms.Set("ololo", "shimba")
	assert.Error(t, err)
	assert.True(t, called)

	called = false
	ms = Store{
		Values: make(map[string]interface{}),
		MockSet: func(s Store, id string, sess interface{}) error {
			called = true
			s.Values[id] = sess
			return nil
		},
	}
	err = ms.Set("ololo", "shimba")
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"ololo": "shimba"}, ms.Values)
	assert.True(t, called)
}

func TestStoreDelete(t *testing.T) {
	called := false
	ms := Store{
		MockDelete: func(_ Store, _ string) error {
			called = true
			return nil
		},
	}
	err := ms.Delete("ololo")
	assert.Nil(t, err)
	assert.True(t, called)

	called = false
	ms = Store{
		MockDelete: func(_ Store, _ string) error {
			called = true
			return errors.New("unknown error")
		},
	}
	err = ms.Delete("ololo")
	assert.Error(t, err)
	assert.True(t, called)

	called = false
	ms = Store{
		Values: map[string]interface{}{"ololo": "shimba"},
		MockDelete: func(s Store, id string) error {
			called = true
			delete(s.Values, id)
			return nil
		},
	}
	err = ms.Delete("ololo")
	assert.Nil(t, err)
	assert.True(t, called)
	assert.Len(t, ms.Values, 0)
}

func TestStoreCount(t *testing.T) {
	expectCount := rand.Intn(1000000)

	called := false
	ms := Store{
		MockCount: func(_ Store) int {
			called = true
			return expectCount
		},
	}
	count := ms.Count()
	assert.Equal(t, expectCount, count)
	assert.True(t, called)

	called = false
	ms = Store{
		Values: map[string]interface{}{"ololo": "shimba"},
		MockCount: func(s Store) int {
			called = true
			return len(s.Values)
		},
	}
	count = ms.Count()
	assert.Equal(t, 1, count)
	assert.True(t, called)
}

func TestStoreVisitAll(t *testing.T) {
	ms := Store{
		Values: map[string]interface{}{
			"ololo":   "shimba",
			"trololo": "shimba",
		},
	}

	var expectedIds []string
	for id := range ms.Values {
		expectedIds = append(expectedIds, id)
	}

	var visitedIds []string
	ms.VisitAll(func(id string, _ interface{}) bool {
		visitedIds = append(visitedIds, id)
		return true
	})

	sort.StringSlice(expectedIds).Sort()
	sort.StringSlice(visitedIds).Sort()
	assert.Equal(t, expectedIds, visitedIds)
}
