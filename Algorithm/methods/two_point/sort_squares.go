package main

import "fmt"

func sortSquares(arr []int) []int {
	res := make([]int, len(arr))
	i := len(arr) - 1
	for l, r := 0, len(arr)-1; l <= r; i-- {
		if arr[l]*arr[l] > arr[r]*arr[r] {
			res[i] = arr[l] * arr[l]
			l++
		} else {
			res[i] = arr[r] * arr[r]
			r--
		}
	}
	return res
}

func main() {
	arr := []int{-11, -10, -2, 0, 1, 5, 6, 11}
	fmt.Println(sortSquares(arr))
}
