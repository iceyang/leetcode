package main

import (
	"sort"
	"strings"
)

func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// func contains(arr []string, str string) bool {
// 	for _, v := range arr {
// 		if v == str {
// 			return true
// 		}
// 	}
// 	return false
// }

func groupAnagrams(strs []string) [][]string {
	hash := map[string][]string{}
	for _, str := range strs {
		key := sortString(str)
		_, ok := hash[key]
		if !ok {
			arr := []string{}
			hash[key] = arr
		}
		// if !contains(hash[key], str) {
		// 	hash[key] = append(hash[key], str)
		// }
		hash[key] = append(hash[key], str)
	}
	result := make([][]string, 0, len(hash))
	for key, _ := range hash {
		result = append(result, hash[key])
	}
	return result
}
