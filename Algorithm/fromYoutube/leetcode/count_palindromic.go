package main

import "fmt"

func main() {
	s := "bbcbaba"
	fmt.Println(countPalindromicSubsequence(s))
}
func countPalindromicSubsequence(s string) int {
	mask := [26]int{}
	for i := range s {
		mask[s[i]-'a']++
	}

	res := 0
	for i := range 26 {
		if mask[i] >= 2 {
			for j := range 26 {
				if j != i && mask[j] > 0 {
					res += 1
				}
			}
			if mask[i] > 2 {
				res++
			}
		}
	}
	return res
}
