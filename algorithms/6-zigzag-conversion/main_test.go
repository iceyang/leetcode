package main

import "testing"

func TestConvert(t *testing.T) {
	if convert("LEETCODEISHIRING", 4) != "LDREOEIIECIHNTSG" {
		t.Fatal()
	}
	if convert("LEETCODEISHIRING", 3) != "LCIRETOESIIGEDHN" {
		t.Fatal()
	}
	convert("PAYPALISHIRING", 9)
}
