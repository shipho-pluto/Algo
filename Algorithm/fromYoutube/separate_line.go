package main

import "fmt"

type Point struct {
	x, y int
}

func main() {
	arr := []Point{{0, 0}, {0, 1}, {1, 1}, {2, 2}, {3, 1}, {4, 1}, {4, 0}}
	fmt.Println(separateLine(arr))
}

func separateLine(arr []Point) bool {
	st := make(map[Point]int)
	n := len(arr)
	maxX, minX := arr[0].x, arr[0].x
	for i := range n {
		st[arr[i]]++
		minX = min(minX, arr[i].x)
		maxX = max(maxX, arr[i].x)
	}

	lineX := maxX + minX

	for k := range st {
		wantPoint := Point{lineX - k.x, k.y}
		if st[k] != st[wantPoint] {
			return false
		}
	}
	return true
}
