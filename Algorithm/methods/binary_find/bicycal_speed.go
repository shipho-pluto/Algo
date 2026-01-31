package main

import "fmt"

func main() {
	x, s := []float64{1, 2, 3, 4}, []float64{3, 3, 1, 2}
	fmt.Println(timeWhenFirstNearLast(0, x[len(x)-1]-x[0], 0.1, s, x))
}

func timeWhenFirstNearLast(l, r, eps float64, s, x []float64) float64 {
	for l+eps < r {
		m := (l + r) / 2
		if fact(m, eps, s, x) {
			r = m
		} else {
			l = m
		}
	}
	return l
}

func dist(m float64, s, x []float64) float64 {
	minPos, maxPos := x[0]+s[0]*m, x[0]+s[0]*m
	for i := 1; i < len(x); i++ {
		nowPos := x[i] + s[i]*m
		minPos = min(nowPos, minPos)
		maxPos = min(maxPos, nowPos)
	}
	return maxPos - minPos
}

func fact(t, eps float64, s, x []float64) bool {
	return dist(t+eps, s, x) >= dist(t, s, x)
}
