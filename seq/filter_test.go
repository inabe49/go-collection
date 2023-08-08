package seq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilter(t *testing.T) {
	assert.Equal(t, []int32{}, Filter[int32]([]int32{}, func(i int32) bool { return true }))
	assert.Equal(t, []int32{}, Filter[int32]([]int32{}, func(i int32) bool { return false }))

	assert.Equal(t, []int32{}, Filter[int32]([]int32{1, 2, 3}, func(i int32) bool { return false }))
	assert.Equal(t, []int32{1, 2, 3}, Filter[int32]([]int32{1, 2, 3}, func(i int32) bool { return true }))

	assert.Equal(t, []int32{2, 4}, Filter[int32]([]int32{1, 2, 3, 4, 5}, func(i int32) bool { return i%2 == 0 }))
}
