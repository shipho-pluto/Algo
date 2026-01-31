package main

import "fmt"

func maxSubString(s string) int {
	mask := [26]int{}
	n := len(s)
	l := 0
	res := 0
	for r := range n {
		if mask[s[r]-'a'] == 1 {
			for ; l < r; l++ {
				if mask[s[r]-'a'] == 0 {
					break
				}
				mask[s[l]-'a']--
			}
		}
		res = max(res, r-l+1)
		mask[s[r]-'a']++
	}

	return res
}

func maxSubString2(s string) int {
	res := 0
	n := len(s)
	l := 0
	mask := [26]int{}
	for r := 0; r < n; r++ {
		l = max(l, mask[s[r]-'a'])
		res = max(res, r-l+1)
		mask[s[r]-'a'] = r + 1
	}

	return res
}

func main() {
	s := "abcabcaartyuiaa"
	fmt.Println(maxSubString(s))
	fmt.Println(maxSubString2(s))
}
