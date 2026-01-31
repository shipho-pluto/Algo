package main

import "fmt"

func foo8(arr []int) int {
	nax := 0
	n := len(arr)
	l, r := 0, n-1
	for l != r {
		nax = max(min(arr[l], arr[r])*(r-l), nax)
		if arr[r] < arr[l] {
			r--
		} else {
			l++
		}
	}
	return nax
}

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(foo8(arr))
}
