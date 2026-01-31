package main

import "fmt"

type TreeIn2 struct {
	left  *TreeIn2
	right *TreeIn2
	value int
}

func IfTreeBFound(t *TreeIn2, min, max int) bool {
	if t != nil {
		if t.value < min || t.value > max {
			return false
		}
		minL := min
		minR := t.value
		maxL := t.value
		maxR := max

		return IfTreeBFound(t.left, minL, maxL) && IfTreeBFound(t.right, minR, maxR)
	}
	return true
}

func Max(t *TreeIn2, preMax int) int {
	if t != nil {
		preMax = Max(t.right, t.value)
	}
	return preMax
}

func Min(t *TreeIn2, preMin int) int {
	if t != nil {
		preMin = Max(t.left, t.value)
	}
	return preMin
}

func main() {
	t := TreeIn2{&TreeIn2{nil, nil, 25}, &TreeIn2{nil, nil, 75}, 50}
	fmt.Println(IfTreeBFound(&t, Min(&t, t.value), Max(&t, t.value)))
}
