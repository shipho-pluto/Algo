package main

import "fmt"

func main() {
	cntPeople := 9
	cntParents := 2
	k := 0.3
	fmt.Println(cntParentsMoreK(cntPeople, cntParents, k))
}

func cntParentsMoreK(people int, parents int, k float64) any {
	return lbsForCntParentsMoreK(0, people, people, parents, k)
}

func lbsForCntParentsMoreK(l int, r int, n int, p int, k float64) int {
	for l < r {
		m := (l + r) / 2
		if float64(p+m) < k*float64(n+m) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
