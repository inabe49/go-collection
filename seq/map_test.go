package seq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAssertEqual(t *testing.T) {
	assert.NotEqual(t, []int32{}, []int32{1, 2})

	slice := make([]int32, 10)
	assert.NotEqual(t, []int32{}, slice)
}

func TestMap(t *testing.T) {
	assert.Equal(t, []int32{}, Map[int32, int32]([]int32{}, func(item int32) int32 { return item }))
	assert.Equal(t, []int32{1, 4, 9}, Map[int32, int32]([]int32{1, 2, 3}, func(item int32) int32 { return item * item }))
}

func TestFlatMap(t *testing.T) {
	assert.Equal(t, []int32{}, FlatMap[int32, int32]([]int32{}, func(item int32) []int32 { return []int32{item} }))
	assert.Equal(t, []int32{1, 2, 3}, FlatMap[int32, int32]([]int32{1, 2, 3}, func(item int32) []int32 { return []int32{item} }))

	assert.Equal(t, []int32{}, FlatMap[int32, int32]([]int32{}, func(item int32) []int32 { return []int32{item, item} }))
	assert.Equal(t, []int32{1, 1, 2, 2, 3, 3}, FlatMap[int32, int32]([]int32{1, 2, 3}, func(item int32) []int32 { return []int32{item, item} }))
}

func TestDrop(t *testing.T) {
	assert.Equal(t, []int32{}, Drop[int32]([]int32{}, 0))
	assert.Equal(t, []int32{}, Drop[int32]([]int32{}, 1))
	assert.Equal(t, []int32{}, Drop[int32]([]int32{}, 2))

	assert.Equal(t, []int32{}, Drop[int32]([]int32{1, 2}, 2))
	assert.Equal(t, []int32{}, Drop[int32]([]int32{1, 2, 3, 4, 5}, 5))
	assert.Equal(t, []int32{}, Drop[int32]([]int32{1, 2, 3, 4, 5}, 6))

	assert.Equal(t, []int32{3, 4, 5}, Drop[int32]([]int32{1, 2, 3, 4, 5}, 2))
}

func TestDropRight(t *testing.T) {
	assert.Equal(t, []int32{}, DropRight[int32]([]int32{}, 0))
	assert.Equal(t, []int32{}, DropRight[int32]([]int32{}, 1))
	assert.Equal(t, []int32{}, DropRight[int32]([]int32{}, 2))

	assert.Equal(t, []int32{}, DropRight[int32]([]int32{1, 2}, 2))
	assert.Equal(t, []int32{}, DropRight[int32]([]int32{1, 2, 3, 4, 5}, 5))
	assert.Equal(t, []int32{}, DropRight[int32]([]int32{1, 2, 3, 4, 5}, 6))

	assert.Equal(t, []int32{1, 2, 3}, DropRight[int32]([]int32{1, 2, 3, 4, 5}, 2))
}

func TestDropWhile(t *testing.T) {
	assert.Equal(t, []int32{}, DropWhile[int32]([]int32{}, func(i int32) bool { return true }))
	assert.Equal(t, []int32{}, DropWhile[int32]([]int32{1}, func(i int32) bool { return true }))
	assert.Equal(t, []int32{}, DropWhile[int32]([]int32{1, 2, 3, 4, 5}, func(i int32) bool { return true }))

	assert.Equal(t, []int32{}, DropWhile[int32]([]int32{}, func(i int32) bool { return false }))
	assert.Equal(t, []int32{1, 2, 3, 4, 5}, DropWhile[int32]([]int32{1, 2, 3, 4, 5}, func(i int32) bool { return false }))

	assert.Equal(t, []int32{3, 4, 5}, DropWhile[int32]([]int32{1, 2, 3, 4, 5}, func(i int32) bool { return i < 3 }))
}
