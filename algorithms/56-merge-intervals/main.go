package main

// 先排序
func quickSort(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	guard := intervals[0]
	i := 0
	j := len(intervals) - 1
	for i < j {
		for intervals[j][0] >= guard[0] && i < j {
			j--
		}
		intervals[i] = intervals[j]
		for intervals[i][0] < guard[0] && i < j {
			i++
		}
		intervals[j] = intervals[i]
	}
	intervals[i] = guard
	quickSort(intervals[:i])
	quickSort(intervals[i+1:])
	return intervals
}

func merge(intervals [][]int) [][]int {
	l := len(intervals)
	res := [][]int{}
	if l == 0 {
		return res
	}
	sortedIntervals := quickSort(intervals)
	for i := 0; i < l-1; i++ {
		if sortedIntervals[i][1] >= sortedIntervals[i+1][0] {
			sortedIntervals[i+1][0] = sortedIntervals[i][0]
			if sortedIntervals[i][1] >= sortedIntervals[i+1][1] {
				sortedIntervals[i+1][1] = sortedIntervals[i][1]
			}
		} else {
			res = append(res, sortedIntervals[i])
		}
	}
	res = append(res, sortedIntervals[l-1])
	return res
}
