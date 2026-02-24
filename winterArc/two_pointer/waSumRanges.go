package main

import "fmt"

func main() {
	arr := []int{-1, 0, 2, 3, 7, 8, 9, 11, 12, 15, 16}
	fmt.Println(sumRanges(arr))
}

func sumRanges(arr []int) []int {
	var res []int

	n := len(arr)
	r := 1
	sum := arr[r-1]
	for ; r < n; r++ {
		if arr[r]-arr[r-1] == 1 {
			sum += arr[r]
		} else {
			res = append(res, sum)
			sum = arr[r]
		}
	}

	if arr[r-1]-arr[r-2] == 1 {
		res = append(res, sum)
	} else {
		res = append(res, arr[r-1])
	}

	return res
}
