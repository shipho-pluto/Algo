package main

import "fmt"

func main() {
	s := "abcabcabcqweqwerqrr"
	k := 2
	fmt.Println(maxLenRepeatLessK(s, k))
}

func maxLenRepeatLessK(s string, k int) int {
	n := len(s)
	mask := make([]int, 26)
	res := 0
	for l, r := 0, 0; r < n; r++ {
		if mask[s[r]-'a'] == k {
			res = max(res, r-l)
			for ; l < r; l++ {
				if mask[s[r]-'a'] < k {
					break
				}
				mask[s[l]-'a']--
			}
		}
		mask[s[r]-'a']++
	}
	return res
}
