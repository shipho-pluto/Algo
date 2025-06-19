package main

import "fmt"

func MaxLenStr(s string) int {
	mask := [26]int{}
	n := len(s)
	maxLen := 0
	sum := 0
	for i := 0; i < n; i++ {
		if mask[s[i]-'a'] != 0 {
			maxLen = max(maxLen, sum)
			sum = i - mask[s[i]-'a']
		}
		sum++
		mask[s[i]-'a'] = i + 1
	}

	return maxLen
}

func main() {
	s := "saabcabcbb"
	fmt.Println(MaxLenStr(s))
}
