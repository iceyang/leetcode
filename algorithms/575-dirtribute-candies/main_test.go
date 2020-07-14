package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistributeCandies(t *testing.T) {
	assert.Equal(t, 3, distributeCandies([]int{1, 1, 2, 2, 3, 3}))
	assert.Equal(t, 2, distributeCandies([]int{1, 1, 1, 2, 2, 2}))
	assert.Equal(t, 0, distributeCandies([]int{}))
	assert.Equal(t, 1, distributeCandies([]int{1, 1, 1, 1}))
}
