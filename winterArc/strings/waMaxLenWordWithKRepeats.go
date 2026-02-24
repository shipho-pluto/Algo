package main

import "fmt"

func main() {
	s := "abcabcabcqweqwergrr"
	k := 2
	fmt.Println(maxLenRepeatLessK(s, k))
}

func maxLenRepeatLessK(s string, k int) int {
	mask := [26]int{}
	maxL := 0
	for l, r := 0, 0; r < len(s); r++ {
		mask[s[r]-'a']++
		if mask[s[r]-'a'] > k {
			for ; l < r; l++ {
				mask[s[l]-'a']--
				if mask[s[l]-'a'] == k {
					break
				}
			}
		}
		maxL = max(maxL, r-l)
	}

	return maxL
}
