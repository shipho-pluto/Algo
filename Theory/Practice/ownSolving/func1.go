package main

import "fmt"

func sum1(arr []int) int {
	s := 0
	for _, v := range arr {
		s += v
	}
	return s
}

func Is1(q bool) int {
	if q {
		return 1
	}
	return 0
}

func Del1El(arr []int) int {
	n, s := len(arr), sum1(arr)
	if n == s || n == s+1 {
		return n - 1
	}
	res := 0
	r, c := 0, Is1(arr[0] == 0)
	for l := 0; l < n; l++ {
		for ; r+1 < n && c+Is1(arr[r+1] == 0) < 2; r++ {
			c += Is1(arr[r+1] == 0)
		}
		res = max(res, r-l)
		c -= Is1(arr[l] == 0)
	}
	return res
}

func main() {
	arr := []int{1, 0, 1, 1, 1, 0, 1, 0, 0, 0, 1, 1, 1, 1, 1, 0, 1}
	fmt.Println(Del1El(arr))
}
