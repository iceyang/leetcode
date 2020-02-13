package main

func trap(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	top := 0
	for i, h := range height {
		if h > height[top] {
			top = i
		}
	}
	total, begin := 0, 0
	for i := 1; i < top; i++ {
		if height[begin] > height[i] {
			total += height[begin] - height[i]
		} else if height[begin] <= height[i] {
			begin = i
		}
	}
	begin = length - 1
	for j := length - 2; j > top; j-- {
		if height[begin] > height[j] {
			total += height[begin] - height[j]
		} else if height[begin] <= height[j] {
			begin = j
		}
	}
	return total
}
