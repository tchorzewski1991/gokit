package collection_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tchorzewski1991/gokit/collection"
	"github.com/tchorzewski1991/gokit/collection/matchers"
)

func TestFind(t *testing.T) {
	testFindString(t)
	testFindInt(t)
	testFindCustomInt(t)
	testFindFloat(t)
	testFindBool(t)
}

func testFindString(t *testing.T) {
	str, ok := collection.Find([]string{"a", "b", "c"}, matchers.Str("b"))
	assert.True(t, ok)
	assert.Equal(t, "b", str)

	str, ok = collection.Find([]string{"a", "b", "c"}, matchers.Str("d"))
	assert.False(t, ok)
	assert.Equal(t, "", str)
}

func testFindInt(t *testing.T) {
	i, ok := collection.Find([]int{1, 2, 3}, matchers.Int(2))
	assert.True(t, ok)
	assert.Equal(t, 2, i)

	i, ok = collection.Find([]int{1, 2, 3}, matchers.Int(4))
	assert.False(t, ok)
	assert.Equal(t, 0, i)
}

func testFindCustomInt(t *testing.T) {
	type customInt int

	i, ok := collection.Find([]customInt{1, 2, 3}, matchers.Int(customInt(1)))
	assert.True(t, ok)
	assert.Equal(t, customInt(1), i)

	i, ok = collection.Find([]customInt{1, 2, 3}, matchers.Int(customInt(4)))
	assert.False(t, ok)
	assert.Equal(t, customInt(0), i)
}

func testFindFloat(t *testing.T) {
	f, ok := collection.Find([]float64{1.1, 2.2, 3.3}, matchers.Float(2.2))
	assert.True(t, ok)
	assert.Equal(t, 2.2, f)

	f, ok = collection.Find([]float64{1.1, 2.2, 3.3}, matchers.Float(2.3))
	assert.False(t, ok)
	assert.Equal(t, float64(0), f)
}

func testFindBool(t *testing.T) {
	res, ok := collection.Find([]bool{true, false}, matchers.Bool(false))
	assert.True(t, ok)
	assert.Equal(t, false, res)

	res, ok = collection.Find([]bool{false, false}, matchers.Bool(true))
	assert.False(t, ok)
	assert.Equal(t, false, res)
}
