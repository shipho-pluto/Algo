package main

import "fmt"

func main() {
	places := []int{2, 5, 7, 11, 15, 20}
	n := 3
	fmt.Println(maxDistanceBetweenEl(places, n))
}

func maxDistanceBetweenEl(places []int, n int) any {
	return rbsForMDBEl(1, places[len(places)-1]-places[0], places, n)
}

func rbsForMDBEl(l int, r int, p []int, n int) int {
	for l < r {
		m := (l + r + 1) / 2
		if checkForMDBEl(m, p, n) {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}

func checkForMDBEl(m int, p []int, n int) bool {
	cnt := 1
	for l, r := 0, 1; r < len(p); r++ {
		if p[r]-p[l] >= m {
			cnt++
			l = r
		}
	}

	return cnt < n
}
