package main

/**
 * 动态规划 - 自下向上
 */
var INT_MAX int = 1<<32 - 1

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	if amount <= 0 {
		return 0
	}
	res := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		res[i] = INT_MAX
		for _, coin := range coins {
			pre := i - coin
			if pre < 0 {
				continue
			}
			res[i] = min(res[i], res[pre]+1)
		}
	}
	if res[amount] == INT_MAX {
		return -1
	}
	return res[amount]
}

// func coinChange(coins []int, amount int) int {
// 	if amount <= 0 {
// 		return 0
// 	}
// 	res := make([]int, amount+1)
// 	for i := 1; i <= amount; i++ {
// 		res[i] = INT_MAX
// 		minOne := 0
// 		for _, coin := range coins {
// 			pre := i - coin
// 			if pre < 0 {
// 				continue
// 			}
// 			res[i] = min(res[i], res[pre]+1)
// 			if pre == 0 {
// 				minOne = 1
// 				continue
// 			}
// 			if pre < 0 || res[pre] == 0 {
// 				continue
// 			}
// 			if minOne == 0 {
// 				minOne = res[pre] + 1
// 				continue
// 			}
// 			if res[pre] < minOne {
// 				minOne = res[pre] + 1
// 			}
// 		}
// 		res[i] = minOne
// 	}
// 	if res[amount] == 0 {
// 		return -1
// 	}
// 	return res[amount]
// }

/**
 * 动态规划 - 自顶向下
 */
// var res []int
//
// func dp(coins []int, amount int) int {
// 	if amount <= 0 {
// 		return 0
// 	}
// 	if res[amount] != 0 {
// 		return res[amount]
// 	}
// 	count := -1
// 	for _, coin := range coins {
// 		if amount-coin < 0 {
// 			continue
// 		}
// 		pre := dp(coins, amount-coin)
// 		if pre == -1 {
// 			continue
// 		}
// 		if count == -1 || pre < count {
// 			count = pre + 1
// 		}
// 	}
// 	res[amount] = count
// 	return count
// }
//
// func coinChange(coins []int, amount int) int {
// 	res = make([]int, amount+1)
// 	return dp(coins, amount)
// }
