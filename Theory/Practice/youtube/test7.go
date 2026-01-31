package main

import "fmt"

type Point7 struct {
	x, y int
}

func abs2(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func HaveVerticalSymLine(arr []Point7) bool {
	n := len(arr)
	l, r := 0, n-1

	if len(arr)%2 == 0 {
		avg := (arr[r].x - arr[l].x) / 2
		for i := 0; i < n/2; i++ {
			if arr[l].y != arr[r].y || abs2(arr[l].x-avg) != abs2(arr[r].x-avg) {
				return false
			}
			l++
			r--
		}
		return true
	} else {
		avg := arr[n/2].x
		for i := 0; i < n/2; i++ {
			if arr[l].y != arr[r].y || abs2(arr[l].x-avg) != abs2(arr[r].x-avg) {
				return false
			}
			l++
			r--
		}
		return true
	}
}

func HaveVerticalSymLine2(arr []Point7) bool {
	pt := make(map[Point7]int)
	n := len(arr)
	maxX, minX := -1000000, 1000000
	for i := 0; i < n; i++ {
		if arr[i].x > maxX {
			maxX = arr[i].x
		}
		if arr[i].x < minX {
			minX = arr[i].x
		}
		pt[arr[i]] += 1
	}

	mid := (maxX - minX) / 2

	for i := 0; i < n; i++ {
		if pt[Point7{mid*2 - arr[i].x, arr[i].y}] != pt[Point7{arr[i].x, arr[i].y}] {
			return false
		}
	}
	return true

}

func main() {
	arr := []Point7{{0, 0}, {0, 1}, {1, 1}, {2, 2}, {3, 1}, {4, 1}, {4, 0}}
	fmt.Println(HaveVerticalSymLine(arr))
	fmt.Println(HaveVerticalSymLine2(arr))
}
