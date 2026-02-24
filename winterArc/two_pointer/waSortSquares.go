package main

import "fmt"

func main() {
	arr := []int{-11, -10, -2, 0, 1, 5, 6, 11}
	fmt.Println(sortSquares(arr))
}

func sortSquares(arr []int) []int {
	res := make([]int, len(arr))
	for l, r, i := 0, len(arr)-1, len(res)-1; i > -1; i-- {
		if arr[l]*arr[l] >= arr[r]*arr[r] {
			res[i] = arr[l] * arr[l]
			l++
		} else {
			res[i] = arr[r] * arr[r]
			r--
		}
	}

	return res
}
