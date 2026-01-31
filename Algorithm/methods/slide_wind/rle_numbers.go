package main

import (
	"fmt"
	"strconv"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 8, 10, 15, 16, 20}
	fmt.Println(rleNumbers(arr))
	arr = []int{-2, -1, 0, 3, 4, 5, 8, 10, 15, 16, 20, 21, 22}
	fmt.Println(rleNumbers(arr))
}

func rleNumbers(arr []int) []string {
	n := len(arr)
	var res []string

	l := 0
	for r := 1; r < n; r++ {
		if arr[r] != arr[r-1]+1 {
			if l != r-1 {
				res = append(res, strconv.Itoa(arr[l])+"->"+strconv.Itoa(arr[r-1]))
			} else {
				res = append(res, strconv.Itoa(arr[l]))
			}
			l = r
		}
	}

	if arr[n-1] != arr[n-2]+1 {
		res = append(res, strconv.Itoa(arr[n-1]))
	} else {
		res = append(res, strconv.Itoa(arr[l])+"->"+strconv.Itoa(arr[n-1]))
	}

	return res
}
