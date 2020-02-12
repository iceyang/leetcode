package main

func trap(height []int) int {
	length := len(height)
	if length == 0 {
		return 0
	}
	total, adding, begin := 0, 0, 0
	for i := 1; i < length; i++ {
		if height[begin] > height[i] && i != length-1 {
			adding = adding + height[begin] - height[i]
			continue
		}
		if height[begin] <= height[i] {
			begin = i
			total += adding
			adding = 0
			continue
		}
		if i == length-1 && begin != i {
			begin += 1
			i = begin
			adding = 0
		}
	}
	return total
}
