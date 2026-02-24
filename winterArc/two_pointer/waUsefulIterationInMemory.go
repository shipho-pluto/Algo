package main

import "fmt"

func main() {
	s := "abcabcaascbs"
	k := 3
	fmt.Println(usefulMemoryIteration(s, k))
}

func usefulMemoryIteration(s string, k int) int {
	res := 0
	for i := k; i < len(s); i++ {
		if s[i] == s[i-k] {
			res++
		}
	}
	return res
}
