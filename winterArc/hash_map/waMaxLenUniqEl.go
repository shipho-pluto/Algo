package main

import "fmt"

func main() {
	s := "abcabcabcdabcqwerty"
	fmt.Println(maxLenUniq(s))
}

func maxLenUniq(s string) int {
	maxL := 0
	st := make(map[uint8]int)
	l := 0
	for r := range s {
		l = max(min(st[s[r]-'a'], r), l)
		st[s[r]-'a'] = r
		maxL = max(maxL, r-l)
	}

	return maxL
}
