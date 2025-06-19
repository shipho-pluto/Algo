package main

import "fmt"

func Road(arr []rune) []rune {
	var res []rune
	mapPoint := make(map[[2]int]bool)
	mapPoint[[2]int{0, 0}] = true
	n := len(arr)
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
		if mapPoint[[2]int{x, y}] {
			cx, cy := x, y
			m := len(res)
			for j := m - 1; j > -1; j-- {
				switch res[j] {
				case 'R':
					cx--
				case 'L':
					cx++
				case 'U':
					cy--
				case 'D':
					cy++
				}
				res = res[:j]
				if cy == y && cx == x {
					break
				}
				mapPoint[[2]int{cx, cy}] = false
			}
		}
		mapPoint[[2]int{x, y}] = true
	}
	return res
}

func main() {
	arr := []rune{'R', 'D', 'L', 'U', 'U', 'L', 'R', 'U'}
	fmt.Println(string(Road(arr)))
}
