package main

/**
 * 利用面积首先于短边的特性，从最大的宽开始，压缩面积计算范围
 * O(n)
 */
func maxArea(height []int) int {
	result := 0
	i := 0
	j := len(height) - 1
	for i < j {
		w := j - i
		value := 0
		if height[i] < height[j] {
			value = w * height[i]
			i++
		} else {
			value = w * height[j]
			j--
		}
		if value > result {
			result = value
		}
	}
	return result
}

/**
 * 暴力解法，性能较低。O(n^2)
 */
func maxArea2(height []int) int {
	result := 0
	length := len(height)
	for x, tall := range height {
		for j := length - 1; j > x; j-- {
			w := j - x
			h := tall
			if height[j] < h {
				h = height[j]
			}
			value := w * h
			if value > result {
				result = value
			}
		}
	}
	return result
}
