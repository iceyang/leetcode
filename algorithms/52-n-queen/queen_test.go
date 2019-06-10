package main

import (
	"testing"
)

func TestTotalNQueens(t *testing.T) {
	if 2 != totalNQueens(4) {
		t.Fail()
	}
	if 40 != totalNQueens(7) {
		t.Fail()
	}
	if 352 != totalNQueens(9) {
		t.Fail()
	}
}
