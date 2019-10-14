package main

import (
	"fmt"
	"strings"
)

type Stack struct {
	node []string
}

func (stack *Stack) push(s string) {
	stack.node = append(stack.node, s)
}

func (stack *Stack) pop() string {
	if len(stack.node) == 0 {
		return ""
	}
	s := stack.node[len(stack.node)-1]
	stack.node = stack.node[:len(stack.node)-1]
	return s
}

func (stack *Stack) size() int {
	return len(stack.node)
}

func simplifyPath(path string) string {
	stack := &Stack{[]string{}}
	paths := strings.Split(path, "/")
	for _, p := range paths {
		if p == "" || p == "." {
			continue
		}
		if p == ".." {
			stack.pop()
		} else {
			stack.push(p)
		}
	}
	fmt.Printf("path: %s, stack: %+v, ", path, stack)
	result := "/" + stack.pop()
	for stack.size() != 0 {
		result = "/" + stack.pop() + result
	}
	fmt.Printf("result: %s\n", result)
	return result
}
