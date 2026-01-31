package main

import "fmt"

func main() {
	arr := []int{0, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 1}
	fmt.Println(maxLenOfOne(arr))
}

func is(q bool) int {
	if q {
		return 1
	}
	return 0
}

func maxLenOfOne(arr []int) int {
	var res int
	n := len(arr)
	cnt := is(arr[0] == 0)
	for l, r := 0, 1; r < n; l++ {
		for r < n && cnt+is(arr[r] == 0) < 2 {
			cnt += is(arr[r] == 0)
			r++
		}
		res = max(res, r-l)
		cnt -= is(arr[l] == 0)
	}

	return res
}
