package seq

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	assert.Equal(t, []int32{}, Map[int32, int32]([]int32{}, func(item int32) int32 { return item }))
	assert.Equal(t, []int32{1, 4, 9}, Map[int32, int32]([]int32{1, 2, 3}, func(item int32) int32 { return item * item }))
}
