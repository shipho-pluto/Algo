package main

import "fmt"

func main() {
	s := "r:-)r:-))r: - )"
	fmt.Println(rleSmile(s))
}

func rleSmile(s string) string {
	cnt := 0
	res := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ':' {
			cnt = 1
		} else if s[i] == '-' && cnt == 1 {
			cnt = 2
		} else if s[i] == ')' && cnt >= 2 {
			cnt += 1
		} else {
			if cnt >= 3 {
				res = res[:len(res)-cnt]
			}
			cnt = 0
		}
		res += string(s[i])

	}

	return res
}
