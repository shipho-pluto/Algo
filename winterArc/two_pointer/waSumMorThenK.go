package main

import "fmt"

func main() {
	arr := []int{-4, -1, 0, 1, 2, 4, 5, 8, 10, 11, 16, 18, 21}
	k := 12
	fmt.Println(difMoreThenK(arr, k))
}

func difMoreThenK(arr []int, k int) int {
	c := 0
	n := len(arr)
	for l, r := 0, 1; r < n && l < n; {
		if -arr[l]+arr[r] <= k {
			r++
		} else {
			l++
			c += n - r
		}
	}
	return c
}
