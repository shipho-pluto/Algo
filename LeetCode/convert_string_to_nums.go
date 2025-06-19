package main

import (
	"fmt"
	"strconv"
)

func ConvertStrToNums(s string) string {
	res := ""
	n := len(s)
	cnt := 1
	for i := 0; i < n-1; i++ {
		if s[i] == s[i+1] {
			cnt++
		} else {
			if cnt > 1 {
				res += string(s[i]) + strconv.Itoa(cnt)
			} else {
				res += string(s[i])
			}
			cnt = 1
		}
	}
	if s[n-1] == s[n-2] {
		res += string(s[n-1]) + strconv.Itoa(cnt)
	} else {
		res += string(s[n-1])
	}

	return res
}

func main() {
	arr := "AAAABBBCCXYZDDDDEEEFFFAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBBBE"
	fmt.Println(ConvertStrToNums(arr))
}
