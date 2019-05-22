package main

import (
	"fmt"
	"testing"
)

func TestMyPow(t *testing.T) {
	fmt.Println(MyPow(2, 4))
	fmt.Println(MyPow(2, 10))
	fmt.Println(MyPow(2, -2))
	fmt.Println(MyPow(2, 31))
	fmt.Println(MyPow(0.00001, 2147483647))
}
