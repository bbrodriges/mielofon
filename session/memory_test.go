package session

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"sort"
	"strconv"
	"testing"
)

func testErrorSeekerOpt(_ StoreSeeker) error {
	return errors.New("bad option")
}

func TestNewMemoryStore(t *testing.T) {
	ms, err := NewMemoryStore()
	assert.NoError(t, err)
	assert.Implements(t, (*Store)(nil), ms)
	assert.Implements(t, (*StoreSeeker)(nil), ms)

	_, err = NewMemoryStore(testErrorSeekerOpt)
	assert.Error(t, err)
}

func TestMemoryStoreGet(t *testing.T) {
	ms, err := NewMemoryStore()
	assert.NoError(t, err)

	id := "2eac4854-fce721f3-b845abba-20d60"
	ms.Map.Store(id, NopSession{})

	sess, err := ms.Get(id)
	assert.NoError(t, err)
	assert.IsType(t, NopSession{}, sess)

	_, err = ms.Get("not_exists")
	assert.Equal(t, ErrNotFound, err)
}

func TestMemoryStoreSet(t *testing.T) {
	ms, err := NewMemoryStore()
	assert.NoError(t, err)

	id := "2eac4854-fce721f3-b845abba-20d60"
	err = ms.Set(id, NopSession{})
	assert.NoError(t, err)

	sess, ok := ms.Map.Load(id)
	assert.True(t, ok)
	assert.IsType(t, NopSession{}, sess)
}

func TestMemoryStoreDelete(t *testing.T) {
	ms, err := NewMemoryStore()
	assert.NoError(t, err)
	assert.Equal(t, 0, ms.Count())

	id := "2eac4854-fce721f3-b845abba-20d60"

	ms.Map.Store(id, NopSession{})
	assert.Equal(t, 1, ms.Count())

	err = ms.Delete(id)
	assert.NoError(t, err)
	assert.Equal(t, 0, ms.Count())
}

func TestMemoryStoreCount(t *testing.T) {
	ms, err := NewMemoryStore()
	assert.NoError(t, err)
	assert.Equal(t, 0, ms.Count())

	id := "2eac4854-fce721f3-b845abba-20d60"
	ms.Map.Store(id, NopSession{})
	assert.Equal(t, 1, ms.Count())

	ms.Map.Delete(id)
	assert.Equal(t, 0, ms.Count())
}

func TestMemoryStoreVisitAll(t *testing.T) {
	ms, err := NewMemoryStore()
	assert.NoError(t, err)

	var ids []string
	for i := 0; i < 5; i++ {
		id := strconv.FormatInt(int64(i), 10)
		err := ms.Set(id, NopSession{})
		assert.NoError(t, err)
		ids = append(ids, id)
	}

	var visitedIds []string
	ms.VisitAll(func(id string, _ interface{}) bool {
		visitedIds = append(visitedIds, id)
		return true
	})

	sort.StringSlice(ids).Sort()
	sort.StringSlice(visitedIds).Sort()
	assert.Equal(t, ids, visitedIds)
}
