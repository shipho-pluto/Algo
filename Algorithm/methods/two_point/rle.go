package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "RAABBCCCRFFFFCCCCCW"
	fmt.Println(rle(s))
}

func rle(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	res := strings.Builder{}
	cnt := 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			cnt++
		} else {
			if cnt == 1 {
				res.WriteString(string(s[i-1]))
			} else {
				res.WriteString(string(s[i-1]) + strconv.Itoa(cnt))
			}
			cnt = 1
		}
	}

	if s[n-1] == s[n-2] {
		res.WriteString(string(s[n-1]) + strconv.Itoa(cnt))
	} else {
		res.WriteString(string(s[n-1]))
	}

	return res.String()
}
