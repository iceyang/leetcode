package main

func distributeCandies(candies []int) int {
	m := make(map[int]bool, len(candies))
	for _, v := range candies {
		m[v] = true
	}
	if len(candies)/2 > len(m) {
		return len(m)
	}
	return len(candies) / 2
}
