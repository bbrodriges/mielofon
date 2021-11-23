package nop

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreGet(t *testing.T) {
	ns := Store{}
	val, err := ns.Get("ololo")
	assert.Nil(t, val)
	assert.Nil(t, err)
}

func TestStoreSet(t *testing.T) {
	ns := Store{}
	assert.Nil(t, ns.Set("ololo", "trololo"))
}

func TestStoreDelete(t *testing.T) {
	ns := Store{}
	assert.Nil(t, ns.Delete("ololo"))
}

func TestStoreCount(t *testing.T) {
	ns := Store{}
	assert.Equal(t, 0, ns.Count())
}

func TestStoreVisitAll(t *testing.T) {
	ns := Store{}

	for i := 0; i < 5; i++ {
		id := strconv.FormatInt(int64(i), 10)
		err := ns.Set(id, "ololo")
		assert.NoError(t, err)
	}

	var visitedIds []string
	ns.VisitAll(func(id string, _ interface{}) bool {
		visitedIds = append(visitedIds, id)
		return true
	})

	assert.Len(t, visitedIds, 0)
}
