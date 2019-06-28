package main

import (
	"fmt"
	"testing"
)

func TestRestoreIpAddresses(t *testing.T) {
	fmt.Println(restoreIpAddresses("25525511135"))
	fmt.Println(restoreIpAddresses("127111"))
	fmt.Println(restoreIpAddresses("127001"))
	fmt.Println(restoreIpAddresses("127244023"))
	fmt.Println(restoreIpAddresses("0000"))
	fmt.Println(restoreIpAddresses("1000000"))
	fmt.Println(restoreIpAddresses("00000"))
}
