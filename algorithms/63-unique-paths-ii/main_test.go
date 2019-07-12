package main

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	if uniquePathsWithObstacles([][]int{
		[]int{0, 0, 0},
		[]int{0, 1, 0},
		[]int{0, 0, 0},
	}) != 2 {
		t.Fatal()
	}
	if uniquePathsWithObstacles([][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
	}) != 10 {
		t.Fatal()
	}
	if uniquePathsWithObstacles([][]int{
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 0},
		[]int{0, 0, 0, 1},
	}) != 0 {
		t.Fatal()
	}
	if uniquePathsWithObstacles([][]int{
		[]int{0, 0, 1, 0},
		[]int{0, 0, 0, 1},
		[]int{0, 0, 0, 0},
	}) != 5 {
		t.Fatal()
	}
	if uniquePathsWithObstacles([][]int{}) != 0 {
		t.Fatal()
	}
	if uniquePathsWithObstacles([][]int{
		[]int{1},
	}) != 0 {
		t.Fatal()
	}
	if uniquePathsWithObstacles([][]int{
		[]int{0},
	}) != 1 {
		t.Fatal()
	}
	if uniquePathsWithObstacles([][]int{
		[]int{1, 0},
	}) != 0 {
		t.Fatal()
	}
	if uniquePathsWithObstacles([][]int{
		[]int{1},
		[]int{0},
	}) != 0 {
		t.Fatal()
	}
	if uniquePathsWithObstacles([][]int{
		[]int{1},
		[]int{1},
	}) != 0 {
		t.Fatal()
	}
	if uniquePathsWithObstacles([][]int{
		[]int{1, 1},
	}) != 0 {
		t.Fatal()
	}
	if uniquePathsWithObstacles([][]int{
		[]int{0, 0},
		[]int{1, 1},
		[]int{0, 0},
	}) != 0 {
		t.Fatal()
	}
}
