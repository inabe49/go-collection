package seq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExists(t *testing.T) {
	assert.False(t, Exists([]int32{}, func(i int32) bool { return i == 3 }))
	assert.False(t, Exists([]int32{}, func(i int32) bool { return true }))
	assert.False(t, Exists([]int32{}, func(i int32) bool { return false }))

	assert.True(t, Exists([]int32{1, 2, 3, 4, 5}, func(i int32) bool { return i == 3 }))
	assert.False(t, Exists([]int32{1, 2, 3, 4, 5}, func(i int32) bool { return i == 100 }))
}
