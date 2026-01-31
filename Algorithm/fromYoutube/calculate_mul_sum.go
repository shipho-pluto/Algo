package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "1*2*3*4*5"
	fmt.Println(calculate(s))
}

func calculate(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	l := 0
	for ; l < n; l++ {
		if s[l] == '*' || s[l] == '+' {
			break
		}
	}

	sum, err := strconv.Atoi(s[:l])
	l++
	if err != nil {
		panic(err)
	}
	prev := sum

	for r := l; r < n; l, r = r+1, r+1 {
		for ; r < n; r++ {
			if s[r] == '+' || s[r] == '*' {
				break
			}
		}
		post, err := strconv.Atoi(s[l:r])
		if err != nil {
			panic(err)
		}

		if s[l-1] == '+' {
			sum += post
			prev = post
		} else if s[l-1] == '*' {
			sum -= prev
			sum += post * prev
			prev *= post
		}
	}

	return sum
}
