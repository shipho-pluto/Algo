package main

import "fmt"

func main() {
	arr := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	fmt.Println(longestOnes(arr, 2))
}

func is(q bool) int {
	if q {
		return 1
	}
	return 0
}

func longestOnes(arr []int, k int) int {
	var cnt, res int
	for l, r := 0, 0; r < len(arr); r++ {
		cnt += is(arr[r] == 0)

		for ; cnt > k; l++ {
			cnt -= is(arr[l] == 0)
		}

		res = max(res, r-l+1)
	}
	return res
}
