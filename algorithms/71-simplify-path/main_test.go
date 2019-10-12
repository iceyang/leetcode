package main

import "testing"

func TestSimplifyPath(t *testing.T) {
	if "/home" != simplifyPath("/home/") {
		t.Fatal("the input '/home/' should be: '/home'")
	}
	if "/" != simplifyPath("/../") {
		t.Fatal("the input '/../' should be: '/'")
	}
	if "/home/foo" != simplifyPath("/home//foo/") {
		t.Fatal("the input '/home//foo/' should be: '/home/foo'")
	}
	if "/c" != simplifyPath("/a/./b/../../c/") {
		t.Fatal("the input '/a/./b/../../c/' should be: '/c'")
	}
	if "/c" != simplifyPath("/a/../../b/../c//.//") {
		t.Fatal("the input '/a/../../b/../c//.//' should be: '/c'")
	}
	if "/a/b/c" != simplifyPath("/a//b////c/d//././/..") {
		t.Fatal("the input '/a//b////c/d//././/..' should be: '/a/b/c'")
	}
}
