package main

import "fmt"

func main() {
	arr := []int{-1, 4, 0, 2, 19, 8, 0, -1}
	l := []int{0, 4, 5}
	r := []int{2, 6, 7}
	fmt.Println(changeArrayByOne(arr, l, r))
}

func changeArrayByOne(arr []int, l, r []int) []int {
	n := len(arr)
	pref := make([]int, n+1)
	li, ri := 0, 0
	for i := range n {
		if li < len(l) && i == l[li] {
			li++
			pref[i+1]++
		} else if ri < len(r) && i == r[ri] {
			pref[i+1]--
			ri++
		}
		pref[i+1] += pref[i]
	}

	for i := 0; i < n; i++ {
		arr[i] += pref[i+1]
	}

	return arr
}
