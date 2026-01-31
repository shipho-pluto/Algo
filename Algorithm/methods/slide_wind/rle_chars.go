package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "yyxxxxxxxxxxxxyy"
	fmt.Println(rleChars(s))
}

func rleChars(s string) []string {
	var res []string
	n := len(s)
	cnt := 1
	for i := 1; i < n; i, cnt = i+1, cnt+1 {
		if s[i] != s[i-1] {
			if cnt == 1 {
				res = append(res, string(s[i-1]))
			} else if cnt > 9 {
				res = append(res, string(s[i-1]))
				strCnt := strconv.Itoa(cnt)
				for j := range strCnt {
					res = append(res, string(strCnt[j]))
				}
			} else {
				res = append(res, string(s[i-1]))
				res = append(res, strconv.Itoa(cnt))
			}
			cnt = 0
		}
	}
	if s[n-1] == s[n-2] {
		if cnt > 9 {
			res = append(res, string(s[n-1]))
			strCnt := strconv.Itoa(cnt)
			for j := range strCnt {
				res = append(res, string(strCnt[j]))
			}
		} else {
			res = append(res, string(s[n-1]))
			res = append(res, strconv.Itoa(cnt))
		}
	} else {
		res = append(res, string(s[n-1]))
	}
	return res
}
