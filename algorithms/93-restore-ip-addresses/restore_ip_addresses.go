package main

import (
	"strconv"
	"strings"
)

var res []string
var segments = make([]string, 0, 4)

func valid(segment string) bool {
	length := len(segment)
	if length == 0 || length > 3 {
		return false
	}
	if segment[0] == '0' && length != 1 {
		return false
	}
	num, _ := strconv.Atoi(segment)
	if num > 255 {
		return false
	}
	return true
}

func dfs(s string, pre int, dots int) {
	length := len(s)
	for i := pre + 1; i < length; i++ {
		if i-pre > 3 {
			break
		}
		segment := s[pre+1 : i+1]
		if valid(segment) {
			segments = append(segments, segment)
			if dots-1 == 0 {
				segment = s[i+1:]
				if valid(segment) {
					segments = append(segments, segment)
					res = append(res, (strings.Join(segments, ".")))
					segments = segments[:len(segments)-1]
				}
			} else {
				dfs(s, i, dots-1)
			}
			segments = segments[:len(segments)-1]
		}
	}
}

func restoreIpAddresses(s string) []string {
	res = make([]string, 0, 20)
	dfs(s, -1, 3)
	return res
}
