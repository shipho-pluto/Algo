package main

import "fmt"

func main() {
	s := "absabcc"
	fmt.Println(cntCouple(s))
}

func cntCouple(s string) int {
	st := [26]int{}
	r, l := 0, 0
	n := len(s)
	for ; r < n; r++ {
		if st[s[r]-'a'] != 0 {
			break
		}
		st[s[r]-'a']++
	}
	res := r * (r - 1) / 2
	for ; r < n; r++ {
		if st[s[r]-'a'] != 0 {
			for ; l < n; l++ {
				if st[s[r]-'a'] == 0 {
					break
				}
				st[s[l]-'a']--
			}
		}
		res += r - l
		st[s[r]-'a']++
	}
	return res
}
