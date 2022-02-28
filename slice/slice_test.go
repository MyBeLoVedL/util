package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceContain(t *testing.T) {
	source := []int{1, 3, 5, 7, 9}
	assert.True(t, ContainInt(source, 1))
	assert.False(t, ContainInt(source, 2))
}

func TestNotNil(t *testing.T) {
	var a *int
	var b []int
	c := []int{}
	var d interface{} = a
	assert.False(t, NotNil(a))
	assert.False(t, NotNil(b))
	assert.True(t, NotNil(c))
	assert.False(t, NotNil(d))
}
