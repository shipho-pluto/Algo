package main

import "fmt"

func main() {
	s := "abcabsaas"
	k := 3
	fmt.Println(cntUsefulIter(s, k))
}

func cntUsefulIter(s string, k int) int {
	res := 0
	for i := k; i < len(s); i++ {
		if s[i] == s[i-k] {
			res++
		}
	}

	return res
}
