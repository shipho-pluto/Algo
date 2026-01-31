package main

import "fmt"

func deleteSmile(s string) string {
	cnt := 0
	res := ""
	n := len(s)
	for i := 0; i < n; i++ {
		res += string(s[i])
		if s[i] == ':' {
			cnt = 1
		} else if s[i] == '-' && cnt == 1 {
			cnt = 2
		} else if s[i] == ')' && cnt == 2 {
			for ; s[i] == ')'; i++ {
				cnt++
			}
			i--
			res = res[:len(res)-cnt]
		} else {
			cnt = 0
		}
	}

	return res
}

func main() {
	s := "r:-)r:-))r: - )"
	fmt.Println(deleteSmile(s))
}
