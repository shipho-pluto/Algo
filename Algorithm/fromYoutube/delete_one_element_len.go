package main

import "fmt"

func main() {
	arr := []int{0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1}
	fmt.Println(deleteElementLen(arr))
}

func sum(arr []int) int {
	sumArr := 0
	for _, v := range arr {
		sumArr += v
	}
	return sumArr
}

func is(q bool) int {
	if q {
		return 1
	}
	return 0
}

func deleteElementLen(arr []int) int {
	n := len(arr)
	s := sum(arr)
	if n == s || n == s+1 {
		return n - 1
	}
	l, r := 0, 0
	maxLen := 0
	cnt := is(arr[0] == 0)
	for ; l < n; l++ {
		for ; r+1 < n && cnt+is(arr[r+1] == 0) < 2; r, cnt = r+1, cnt+is(arr[r+1] == 0) {
		}
		maxLen = max(maxLen, r-l)
		cnt -= is(arr[l] == 0)
	}

	return maxLen
}
