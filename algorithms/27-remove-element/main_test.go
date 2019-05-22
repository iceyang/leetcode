package main

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 2
	fmt.Printf("old array: %+v\n", nums)
	length := RemoveElement(nums, val)
	fmt.Printf("new array: %+v\n", nums[:length])
	fmt.Printf("length: %d\n", length)
}
