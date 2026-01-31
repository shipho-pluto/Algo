package main

import (
	"fmt"
	"strconv"
)

func ConvertStrToNums(s string) string {
	res := string(s[0])
	n := len(s)
	cnt := 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			if cnt > 1 {
				res += string(s[i-1]) + strconv.Itoa(cnt)
			} else {
				res += string(s[i-1])
			}
			cnt = 1
		}
	}
	res += string(s[n-1])
	if s[n-1] == s[n-2] {
		res += strconv.Itoa(cnt)
	}
	return res
}

func main() {
	arr := "AAAABBBCCXYZDDDDEEEFFFAAAAAABBBBBBBBBBBBBBBBBBBBBBBBBBBB"
	fmt.Println(ConvertStrToNums(arr))
}
