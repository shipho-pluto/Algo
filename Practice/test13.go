package main

import "fmt"

func Road(arr []rune) []rune {
	var r []rune
	n := len(arr)
	x, y := 0, 0
	Map := make(map[[2]int]int)
	Map[[2]int{0, 0}] = 1
	for i := 0; i < n; i++ {
		r = append(r, arr[i])
		switch arr[i] {
		case 'R':
			x++
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		}
		if Map[[2]int{x, y}] != 0 {
			xs, ys := x, y
			m := len(r)
			for j := m - 1; j > -1; j-- {
				switch r[j] {
				case 'R':
					x--
				case 'U':
					y--
				case 'D':
					y++
				case 'L':
					x++
				}
				r = r[:j]
				if x == xs && y == ys {
					break
				}
				Map[[2]int{x, y}] = 0
			}
		}
		Map[[2]int{x, y}] = 1
	}
	return r
}

func main() {
	arr := []rune{'R', 'D', 'L', 'U', 'R', 'D', 'L', 'U', 'R', 'R', 'U', 'L', 'U', 'R', 'D', 'L', 'U'}
	for _, val := range Road(arr) {
		fmt.Print(string(val) + " ")
	}
}
