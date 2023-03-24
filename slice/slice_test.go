package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tchorzewski1991/gokit/slice"
)

func TestFind(t *testing.T) {
	elems := []string{"a", "b", "c"}

	s, ok := slice.Find(elems, func(e string) bool {
		return e == "b"
	})
	assert.True(t, ok)
	assert.Equal(t, s, elems[1])

	s, ok = slice.Find(elems, func(e string) bool {
		return e == "d"
	})
	assert.False(t, ok)
	assert.Equal(t, s, "")
}
