package main

import "fmt"

func main() {
	s := "sdgr:-)a:-))::-):-"
	fmt.Println(rleSmile(s))
}

func rleSmile(s string) string {
	var res string
	n := len(s)
	flag := 0
	for r := 0; r < n; r++ {
		if flag == 0 && s[r] == ':' {
			flag = 1
		} else if flag == 1 && s[r] == '-' {
			flag = 2
		} else if flag >= 2 && s[r] == ')' {
			flag++
		} else if flag >= 2 {
			res = res[:len(res)-flag]
			flag = 0
		}
		res += string(s[r])
	}
	return res
}
