package main

import "fmt"

func main() {
	w, h, n := 25, 16, 5
	fmt.Println(lengthRec(w, h, n))
}

func lengthRec(w int, h int, n int) any {
	return rbsForLengthRec(1, min(w, h), w, h, n)
}

func rbsForLengthRec(l, r, w, h, n int) int {
	for l < r {
		m := (l + r + 1) / 2
		if m*n > w || m*n > h {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}
