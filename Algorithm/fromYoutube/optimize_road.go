package main

import "fmt"

func main() {
	arr := []rune{'R', 'L', 'R', 'U', 'U', 'R', 'D', 'L', 'U', 'R', 'D', 'L', 'R', 'R', 'R', 'U'}
	for _, val := range optimizeRoad(arr) {
		fmt.Println(string(val))
	}
}

func optimizeRoad(arr []rune) []rune {
	n := len(arr)
	points := make(map[[2]int]bool)
	res := make([]rune, 0, n)
	x, y := 0, 0
	for i := 0; i < n; i++ {
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
			for j := len(res) - 1; j > -1; j-- {
				switch res[j] {
				case 'R':
					xc--
				case 'L':
					xc++
				case 'U':
					yc--
				case 'D':
					yc++
				}
				res = res[:j]
				points[[2]int{xc, yc}] = false
				if !points[[2]int{x, y}] {
					break
				}
			}
		}
		points[[2]int{x, y}] = true
	}

	return res
}
