package main

import "fmt"

func main() {
	arr := []int{0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1}
	fmt.Println(maxLength(arr, 1) - 1)
}

func maxLength(arr []int, k int) int {
	res := 0
	cnt := 0
	for l, r := 0, 0; r < len(arr); r++ {
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
