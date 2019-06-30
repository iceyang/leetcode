package main

import (
	"testing"
)

func TestMinSteps(t *testing.T) {
	if minSteps(3) != 3 {
		t.Fail()
	}
	if minSteps(6) != 5 {
		t.Fail()
	}
	if minSteps(10) != 7 {
		t.Fail()
	}
	if minSteps(19) != 19 {
		t.Fail()
	}
	if minSteps(100) != 14 {
		t.Fail()
	}
}
