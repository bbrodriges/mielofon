package session

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"sort"
	"testing"
)

func TestMockStoreGet(t *testing.T) {
	called := false
	ms := MockStore{
		MockGet: func(_ MockStore, _ string) (interface{}, error) {
			called = true
			return NopSession{}, nil
		},
	}
	sess, err := ms.Get("ololo")
	assert.Nil(t, err)
	assert.Equal(t, NopSession{}, sess)
	assert.True(t, called)

	called = false
	ms = MockStore{
		MockGet: func(_ MockStore, _ string) (interface{}, error) {
			called = true
			return nil, errors.New("unknown error")
		},
	}
	sess, err = ms.Get("ololo")
	assert.Nil(t, sess)
	assert.Error(t, err)
	assert.True(t, called)

	called = false
	ms = MockStore{
		Store: map[string]interface{}{"ololo": NopSession{}},
		MockGet: func(s MockStore, id string) (interface{}, error) {
			called = true
			if sess, ok := s.Store[id]; ok {
				return sess, nil
			}
			return nil, errors.New("not found")
		},
	}
	sess, err = ms.Get("ololo")
	assert.Equal(t, NopSession{}, sess)
	assert.NoError(t, err)
	assert.True(t, called)
}

func TestMockStoreSet(t *testing.T) {
	called := false
	ms := MockStore{
		MockSet: func(_ MockStore, _ string, _ interface{}) error {
			called = true
			return nil
		},
	}
	err := ms.Set("ololo", NopSession{})
	assert.Nil(t, err)
	assert.True(t, called)

	called = false
	ms = MockStore{
		MockSet: func(_ MockStore, _ string, _ interface{}) error {
			called = true
			return errors.New("unknown error")
		},
	}
	err = ms.Set("ololo", NopSession{})
	assert.Error(t, err)
	assert.True(t, called)

	called = false
	ms = MockStore{
		Store: make(map[string]interface{}),
		MockSet: func(s MockStore, id string, sess interface{}) error {
			called = true
			s.Store[id] = sess
			return nil
		},
	}
	err = ms.Set("ololo", NopSession{})
	assert.NoError(t, err)
	assert.Equal(t, map[string]interface{}{"ololo": NopSession{}}, ms.Store)
	assert.True(t, called)
}

func TestMockStoreDelete(t *testing.T) {
	called := false
	ms := MockStore{
		MockDelete: func(_ MockStore, _ string) error {
			called = true
			return nil
		},
	}
	err := ms.Delete("ololo")
	assert.Nil(t, err)
	assert.True(t, called)

	called = false
	ms = MockStore{
		MockDelete: func(_ MockStore, _ string) error {
			called = true
			return errors.New("unknown error")
		},
	}
	err = ms.Delete("ololo")
	assert.Error(t, err)
	assert.True(t, called)

	called = false
	ms = MockStore{
		Store: map[string]interface{}{"ololo": NopSession{}},
		MockDelete: func(s MockStore, id string) error {
			called = true
			delete(s.Store, id)
			return nil
		},
	}
	err = ms.Delete("ololo")
	assert.Nil(t, err)
	assert.True(t, called)
	assert.Len(t, ms.Store, 0)
}

func TestMockStoreCount(t *testing.T) {
	expectCount := rand.Intn(1000000)

	called := false
	ms := MockStore{
		MockCount: func(_ MockStore) int {
			called = true
			return expectCount
		},
	}
	count := ms.Count()
	assert.Equal(t, expectCount, count)
	assert.True(t, called)

	called = false
	ms = MockStore{
		Store: map[string]interface{}{"ololo": NopSession{}},
		MockCount: func(s MockStore) int {
			called = true
			return len(s.Store)
		},
	}
	count = ms.Count()
	assert.Equal(t, 1, count)
	assert.True(t, called)
}

func TestMockStoreVisitAll(t *testing.T) {
	ms := MockStore{
		Store: map[string]interface{}{
			"ololo":   NopSession{},
			"trololo": NopSession{},
		},
	}

	var expectedIds []string
	for id := range ms.Store {
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
