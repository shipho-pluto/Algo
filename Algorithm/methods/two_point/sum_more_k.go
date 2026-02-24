package main

import "fmt"

func main() {
	arr := []int{-4, -1, 0, 1, 2, 4, 5, 8, 10, 11, 16, 18, 21}
	k := 12
	fmt.Println(sumMoreK(arr, k))
}

func sumMoreK(arr []int, k int) int {
	n := len(arr)
	res := 0
	for l, r := 0, 0; l < n; l++ {
		for ; r < n; r++ {
			if arr[r]-arr[l] >= k {
				break
			}
		}
		res += n - r
	}

	return res
}
