package main

import (
	"fmt"
	"strconv"
)

func main() {
	ex := "1+2*4+2+1*4*2"
	s := "1*2*3*4*5"
	fmt.Println(calculate(ex))
	fmt.Print(calculate(s))
}

func calculate(ex string) int {
	l := 0
	for ; l < len(ex); l++ {
		if ex[l] == '*' || ex[l] == '+' {
			break
		}
	}

	prev, err := strconv.Atoi(ex[:l])
	if err != nil {
		panic(1)
	}
	sum := prev
	for r := l + 1; r < len(ex); l, r = r, r+1 {
		for ; r < len(ex) && ex[r] != '*' && ex[r] != '+'; r++ {
		}
		cur, err := strconv.Atoi(ex[l+1 : r])
		if err != nil {
			panic(1)
		}

		if ex[l] == '+' {
			sum += cur
			prev = cur
		} else if ex[l] == '*' {
			sum -= prev
			sum += prev * cur
			prev = prev * cur
		}
	}

	return sum
}
