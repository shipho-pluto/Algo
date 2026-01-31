package main

import "fmt"

func main() {
	w, h, n := 25, 16, 5
	fmt.Println(howManyRec(w, h, n))
}

func howManyRec(w int, h int, n int) any {
	return bfHMR(0, min(w, h), w, h, n)
}

func bfHMR(l, r int, w, h, n int) int {
	for l < r {
		m := (l + r + 1) / 2
		if m*n <= h && m*n <= w {
			l = m
		} else {
			r = m - 1
		}
	}
	return l
}
