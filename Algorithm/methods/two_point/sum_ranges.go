package main

import "fmt"

func main() {
	arr := []int{-1, 0, 2, 3, 7, 8, 9, 11, 12, 15, 16, 20}
	fmt.Println(sumRanges(arr))
}

func sumRanges(arr []int) []int {
	n := len(arr)
	if n == 1 {
		return arr
	}
	res := make([]int, 0, n)
	sum := arr[0]
	for i := 1; i < n; i++ {
		if arr[i]-1 == arr[i-1] {
			sum += arr[i]
		} else {
			res = append(res, sum)
			sum = arr[i]
		}
	}

	if arr[n-1]-1 == arr[n-2] {
		res = append(res, sum)
	} else {
		res = append(res, arr[n-1])
	}

	return res
}
