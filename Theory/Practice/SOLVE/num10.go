package main

import "fmt"

func correctWay(arr []rune) []rune {
	points := make(map[[2]int]bool)
	var res []rune
	n := len(arr)
	x, y := 0, 0
	for i := 0; i < n; i++ {
		switch arr[i] {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
		res = append(res, arr[i])
		if points[[2]int{x, y}] == true {
			xc, yc := x, y
			for j := len(res) - 1; j > -1; j-- {
				switch res[j] {
				case 'U':
					yc--
				case 'D':
					yc++
				case 'L':
					xc++
				case 'R':
					xc--
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

func main() {
	way := []rune{'U', 'L', 'D', 'R', 'U', 'R', 'U', 'D', 'R', 'U'}
	cWay := correctWay(way)
	for i := range cWay {
		fmt.Print(string(cWay[i]) + " ")
	}
}
