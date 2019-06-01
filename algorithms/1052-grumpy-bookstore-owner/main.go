package main

/**
 * 核心在于，如何在X连续时间内，转换得到最大的满意度
 * 也就是，我们只关心不满意的情况，忽略满意的情况，
 * 只要得到计算滑动窗口内不满意的最大值是多少，再加上原本的满意值，
 * 自然就能得到最大满意值
 */
func maxSatisfied(customers []int, grumpy []int, X int) int {
	result := 0
	window := 0
	for i := 0; i < X; i++ {
		if grumpy[i] == 1 {
			window += customers[i]
		} else {
			result += customers[i]
		}
	}
	maxWindow := window
	for i := X; i < len(customers); i++ {
		window = window - customers[i-X]*grumpy[i-X]
		if grumpy[i] != 1 {
			result += customers[i]
			continue
		}
		window += customers[i]
		if window > maxWindow {
			maxWindow = window
		}
	}
	result += maxWindow
	return result
}

func maxSatisfied2(customers []int, grumpy []int, X int) int {
	result := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			result += customers[i]
		}
	}
	window := 0
	for i := 0; i < X; i++ {
		if grumpy[i] == 1 {
			window += customers[i]
		}
	}
	maxWindow := window
	for i := X; i < len(customers); i++ {
		window = window - customers[i-X]*grumpy[i-X]
		if grumpy[i] != 1 {
			continue
		}
		window += customers[i]
		if window > maxWindow {
			maxWindow = window
		}
	}
	result += maxWindow
	return result
}

func maxSatisfied1(customers []int, grumpy []int, X int) int {
	result := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 0 {
			result += customers[i]
		}
	}

	// 计算X能得到的最大效益
	startIndex := 0
	i := 0
	window := 0
	for i < X {
		if grumpy[i] == 1 {
			window += customers[i]
		}
		i++
	}
	maxWindow := window
	for ; i < len(customers); i++ {
		window = window - customers[i-X]*grumpy[i-X]
		if grumpy[i] != 1 {
			continue
		}
		window += customers[i]
		if window > maxWindow {
			startIndex = i - X + 1
			maxWindow = window
		}
	}
	for j := 0; j < X; j++ {
		if grumpy[j+startIndex] == 1 {
			result += customers[j+startIndex]
		}
	}
	return result
}
