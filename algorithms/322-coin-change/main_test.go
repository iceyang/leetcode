package main

import "testing"

func TestCoinChange(t *testing.T) {
	t.Log(coinChange([]int{1, 2, 5}, 11) == 3)
	t.Log(coinChange([]int{2}, 3) == -1)
	t.Log(coinChange([]int{5, 11}, 150) == 18)
	t.Log(coinChange([]int{5, 11}, 15) == 3)
	t.Log(coinChange([]int{5, 11}, 0) == 0)
	t.Log(coinChange([]int{}, 5) == -1)
	t.Log(coinChange([]int{}, 0) == 0)
}
