package main

import "fmt"

func main() {
	s := "yxyabcxyx"
	fmt.Println(maxUniqSubstring(s))
}

func maxUniqSubstring(s string) int {
	var res int
	mask := [26]int{}
	post := 0
	for i := range s {
		post = max(post, mask[s[i]-'a'])
		res = max(res, i-post)
		mask[s[i]-'a'] = i
	}
	return res
}
