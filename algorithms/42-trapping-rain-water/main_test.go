package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrap(t *testing.T) {
	assert.Equal(t, 6, trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(t, 0, trap([]int{0, 0, 0, 0}))
	assert.Equal(t, 0, trap([]int{1, 1, 1, 1}))
	assert.Equal(t, 2, trap([]int{1, 0, 0, 1}))
	assert.Equal(t, 2, trap([]int{2, 0, 0, 1}))
	assert.Equal(t, 9, trap([]int{0, 1, 3, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
