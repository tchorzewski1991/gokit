package matchers_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tchorzewski1991/gokit/slice"
	"github.com/tchorzewski1991/gokit/slice/matchers"
)

func TestStr(t *testing.T) {
	str, ok := slice.Find([]string{"a", "b", "c"}, matchers.Str("b"))
	assert.True(t, ok)
	assert.Equal(t, "b", str)

	str, ok = slice.Find([]string{"a", "b", "c"}, matchers.Str("d"))
	assert.False(t, ok)
	assert.Equal(t, "", str)
}

func TestInt(t *testing.T) {
	i, ok := slice.Find([]int{1, 2, 3}, matchers.Int(2))
	assert.True(t, ok)
	assert.Equal(t, 2, i)

	i, ok = slice.Find([]int{1, 2, 3}, matchers.Int(4))
	assert.False(t, ok)
	assert.Equal(t, 0, i)

}

func TestIntLike(t *testing.T) {
	type customInt int

	i, ok := slice.Find([]customInt{1, 2, 3}, matchers.Int(customInt(1)))
	assert.True(t, ok)
	assert.Equal(t, customInt(1), i)

	i, ok = slice.Find([]customInt{1, 2, 3}, matchers.Int(customInt(4)))
	assert.False(t, ok)
	assert.Equal(t, customInt(0), i)
}

func TestFloat(t *testing.T) {
	f, ok := slice.Find([]float64{1.1, 2.2, 3.3}, matchers.Float(2.2))
	assert.True(t, ok)
	assert.Equal(t, 2.2, f)

	f, ok = slice.Find([]float64{1.1, 2.2, 3.3}, matchers.Float(2.3))
	assert.False(t, ok)
	assert.Equal(t, float64(0), f)
}

func TestBool(t *testing.T) {
	res, ok := slice.Find([]bool{true, false}, matchers.Bool(false))
	assert.True(t, ok)
	assert.Equal(t, false, res)

	res, ok = slice.Find([]bool{false, false}, matchers.Bool(true))
	assert.False(t, ok)
	assert.Equal(t, false, res)
}
