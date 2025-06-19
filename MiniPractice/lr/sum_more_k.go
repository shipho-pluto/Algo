package main

import "fmt"

func SumMoreThanK(arr []int, k int) int {
	cnt := 0
	n := len(arr)
	l, r := 0, 1
	for ; l < r; l++ {
		for ; r < n && arr[r]-arr[l] <= k; r++ {
		}
		cnt += n - r
	}
	return cnt
}

func main() {
	arr := []int{-4, -1, 0, 1, 2, 4, 5, 8, 10, 11, 16, 18, 21}
	f := 12
	fmt.Println(SumMoreThanK(arr, f))
}
