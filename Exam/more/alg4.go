package main

import "fmt"

func maxLenRepeatElementLessK(arr string, k int) int {
	n := len(arr)
	mask := [26]int{}
	maxLen := 0
	for l, r := 0, 0; r < n; r++ {
		mask[arr[r]-'a']++
		if mask[arr[r]-'a'] == k {
			for l < n && mask[arr[r]-'a'] == k {
				mask[arr[l]-'a']--
				l++
			}
		}
		maxLen = max(maxLen, r-l)
	}
	return maxLen
}

func main() {
	arr := "abcabcabccabcabcafdddds"
	k := 2
	fmt.Println(maxLenRepeatElementLessK(arr, k))
}
