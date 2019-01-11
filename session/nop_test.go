package session

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestNopStoreGet(t *testing.T) {
	ns := NopStore{}
	val, err := ns.Get("ololo")
	assert.Equal(t, NopSession{}, val)
	assert.Nil(t, err)
}

func TestNopStoreSet(t *testing.T) {
	ns := NopStore{}
	assert.Nil(t, ns.Set("ololo", NopSession{}))
}

func TestNopStoreDelete(t *testing.T) {
	ns := NopStore{}
	assert.Nil(t, ns.Delete("ololo"))
}

func TestNopStoreCount(t *testing.T) {
	ns := NopStore{}
	assert.Equal(t, 0, ns.Count())
}

func TestNopStoreVisitAll(t *testing.T) {
	ns := NopStore{}

	for i := 0; i < 5; i++ {
		id := strconv.FormatInt(int64(i), 10)
		err := ns.Set(id, NopSession{})
		assert.NoError(t, err)
	}

	var visitedIds []string
	ns.VisitAll(func(id string, _ interface{}) bool {
		visitedIds = append(visitedIds, id)
		return true
	})

	assert.Len(t, visitedIds, 0)
}
