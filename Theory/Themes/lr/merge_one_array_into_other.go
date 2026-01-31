package main

import "fmt"

func mergeArrayIntoAnother(a1 []int, n int, a2 []int, m int) {
	l, r := n-1, m-1
	i := n + m - 1
	for l > -1 || r > -1 {
		if l > -1 && r > -1 {
			if a1[l] > a2[r] {
				a1[i] = a1[l]
				l--
				i--
			} else {
				a1[i] = a2[r]
				r--
				i--
			}
		} else if l > -1 {
			a1[i] = a1[l]
			l--
			i--
		} else {
			a1[i] = a2[r]
			r--
			i--
		}
	}
}

func main() {
	n1, n2 := []int{1, 2, 3, 0, 0, 0}, []int{2, 4, 5}
	mergeArrayIntoAnother(n1, 3, n2, 3)
	fmt.Println(n1)
}
