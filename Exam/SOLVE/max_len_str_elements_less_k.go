package main

import (
	"fmt"
)

func maxLenElementLessK(s string, k int) int {
	n := len(s)
	st := make(map[string]int)
	res := 0
	i, j := 0, 0
	for ; j < n; j++ {
		if st[string(s[j])]+1 > k {
			res = max(j-i, res)
			for ; i < j; i++ {
				if st[string(s[j])]+1 == k {
					break
				}
				st[string(s[i])]--
			}
		}
		st[string(s[j])]++
	}
	res = max(j-i, res)

	return res
}

func main() {
	s := "aaadb"
	k := 2

	fmt.Println(maxLenElementLessK(s, k))
}
