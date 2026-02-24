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

	minX := arr[0].x
	maxX := arr[len(arr)-1].x
	for i := range arr {
		st[arr[i]]++
	}

	LineX := minX + maxX
	for k, v := range st {
		needX := LineX - k.x
		if v != st[Point{needX, k.y}] {
			return false
		}
	}

	return true
}
