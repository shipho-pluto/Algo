package main

import "fmt"

type point struct {
	x, y int
}

func IfThereSepLine(arr []point) bool {
	st := map[point]int{}
	minX, maxX := 0, 0
	for i := range arr {
		minX = min(minX, arr[i].x)
		maxX = max(maxX, arr[i].x)
		st[arr[i]]++
	}

	lineX := maxX + minX
	for i := range st {
		wantPoint := point{x: lineX - i.x, y: i.y}
		if st[i] != st[wantPoint] {
			return false
		}
	}
	return true
}

func main() {
	a := []point{{0, 0}, {0, 1}, {1, 1}, {2, 2}, {3, 1}, {4, 1}, {4, 0}}
	fmt.Println(IfThereSepLine(a))
}
