package main

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestReversePairs(t *testing.T) {
	assert.Equal(t, 5, reversePairs([]int{7, 5, 6, 4}))
	assert.Equal(t, 6, reversePairs([]int{4, 3, 2, 1}))
}
