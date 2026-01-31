package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	arr := []point{{0, 0}, {0, 1}, {1, 1}, {2, 2}, {3, 1}, {4, 1}, {4, 0}}
	fmt.Println(existLine(arr))
}

func existLine(arr []point) bool {
	n := len(arr)
	minX, maxX := arr[0].x, arr[0].x
	st := make(map[point]int)
	for i := range n {
		st[arr[i]]++
		minX = min(minX, arr[i].x)
		maxX = max(maxX, arr[i].x)
	}
	lineX := maxX + minX
	for k, v := range st {
		wantPoint := point{x: lineX - k.x, y: k.y}
		if st[wantPoint] != v {
			return false
		}
	}
	return true
}
