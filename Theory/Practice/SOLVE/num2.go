package main

import "fmt"

func sumArr(arr []int) int {
	sum := 0
	for i := range arr {
		sum += arr[i]
	}
	return sum
}

func Is(q bool) int {
	if q {
		return 1
	}
	return 0
}

func DeleteElementMax(arr []int) int {
	n := len(arr)
	if n == sumArr(arr) || n == sumArr(arr)+1 {
		return n - 1
	}
	l, r := 0, 0
	c := Is(arr[0] == 0)
	maxLen := 0
	for l < n {
		for ; r < n-1 && c+Is(arr[r+1] == 0) < 2; func() { r++; c += Is(arr[r] == 0) }() {
		}
		maxLen = max(maxLen, r-l)
		c -= Is(arr[l] == 0)
		l++
	}
	return maxLen
}

func main() {
	a := []int{0, 1, 0, 1, 1, 0, 0, 1, 1, 0, 1, 1, 0, 1, 1, 1}
	fmt.Println(DeleteElementMax(a))
}
