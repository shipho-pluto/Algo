package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "RAABBCCCRFFFFCCCCCW"
	fmt.Println(rle(str))
}

func rle(s string) string {
	newSTR := ""
	l, r := 0, 1
	for r < len(s) {
		if s[r] == s[l] {
			r++
		} else {
			if r-l > 1 {
				newSTR += string(s[l]) + strconv.Itoa(r-l)
			} else {
				newSTR += string(s[l])
			}
			l = r
		}
	}

	if r-l > 2 {
		newSTR += string(s[r-1]) + strconv.Itoa(r-l-1)
	} else {
		newSTR += string(s[r-1])
	}

	return newSTR
}
