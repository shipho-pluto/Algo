package main

import "fmt"

func main() {
	arr := []rune{'R', 'L', 'R', 'U', 'U', 'R', 'D', 'L', 'U', 'R', 'D', 'L', 'R', 'R', 'R', 'U'}
	for _, val := range optimizeRoad(arr) {
		fmt.Print(string(val) + " ")
	}
}

func optimizeRoad(arr []rune) []rune {
	var res []rune
	points := make(map[[2]int]bool)
	points[[2]int{0, 0}] = true
	x, y := 0, 0
	for i := range arr {
		switch arr[i] {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
		res = append(res, arr[i])
		if points[[2]int{x, y}] {
			xc, yc := x, y
			for j := i; j > -1; j-- {
				switch arr[j] {
				case 'R':
					xc--
				case 'L':
					xc++
				case 'U':
					yc--
				case 'D':
					yc++
				}
				res = res[:len(res)-1]
				points[[2]int{xc, yc}] = false
				if x == xc && y == yc {
					break
				}
			}
		}
		points[[2]int{x, y}] = true
	}

	return res
}
