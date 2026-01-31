package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	fmt.Println(longestOnes(arr, 2))
}

func longestOnes(arr []int, k int) int {
	res, cnt := 0, 0
	for r, l := 0, 0; r < len(arr); r++ {
		if arr[r] == 0 {
			cnt++
		}
		for ; cnt > k; l++ {
			if arr[l] == 0 {
				cnt--
			}
		}
		res = max(res, r-l+1)
	}

	return res
}
