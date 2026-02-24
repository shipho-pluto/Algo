package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	k := 4
	replace(a, k)
	fmt.Println(a)
}

func replace(a []int, k int) {
	kc := k
	for l, r := k-1, len(a)-1; r > 0; {
		if kc > 0 {
			a[l], a[r] = a[r], a[l]
			l, r, kc = l-1, r-1, kc-1
		} else {
			l = min(k-1, r-1)
			kc = k
		}
	}
}
