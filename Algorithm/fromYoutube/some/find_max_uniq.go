package main

import "fmt"

func main() {
	s := "pwwkew"
	fmt.Println(maxUniq(s))
}

func maxUniq(s string) any {
	st := [26]int{}
	prev := 0
	res := 0
	for i := range s {
		prev = max(prev, st[s[i]-'a'])
		res = max(res, i-prev)
		st[s[i]-'a'] = i
	}
	return res
}
