package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)

	for l, r := 0, len(s)-1; l < r; func() { l++; r-- }() {
		if s[l] != s[r] {
			return false
		}
	}
	return true
}

func main() {
	p := "aboba"
	fmt.Println(isPalindrome(p))
}
