package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAndSay(t *testing.T) {
	assert.Equal(t, "1", countAndSay(1))
	assert.Equal(t, "111221", countAndSay(5))
	assert.Equal(t, "13211311123113112211", countAndSay(10))
}
