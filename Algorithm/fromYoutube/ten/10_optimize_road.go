package main

import "fmt"

func main() {
	road := []rune{
		'R',
		'R', 'D', 'L', 'U',
		'R', 'D', 'L', 'U',
		'U',
		'U', 'D',
		'R',
		'R', 'L',
	}
	for _, i := range optimizeRoad(road) {
		fmt.Println(string(i))
	}
}

func optimizeRoad(road []rune) []rune {
	var res []rune
	n := len(road)
	points := make(map[[2]int]bool)
	x, y := 0, 0
	for i := 0; i < n; i++ {
		switch road[i] {
		case 'R':
			x++
		case 'L':
			x--
		case 'U':
			y++
		case 'D':
			y--
		}
		res = append(res, road[i])
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
				if x == xc && y == yc {
					break
				}
			}
		}
		points[[2]int{x, y}] = true
	}
	return res
}
