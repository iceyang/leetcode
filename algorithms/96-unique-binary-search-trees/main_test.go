package main

import "testing"

func TestNumTrees(t *testing.T) {
	if numTrees(3) != 5 {
		t.Fatal()
	}
}
