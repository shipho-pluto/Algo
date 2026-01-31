package main

import "fmt"

func main() {
	arr := []int{2, 1, 0, 6}
	fmt.Println(increasingTripleSubsequence(arr))
}

func increasingTripleSubsequence(arr []int) bool {
	n := len(arr)
	for i := 0; i < n; i++ {
		l := i
		for l < n && arr[l] <= arr[i] {
			l++
		}
		if l == n {
			continue
		}
		r := l
		for r < n && arr[r] <= arr[l] {
			r++
		}
		if r != n {
			return true
		}
	}
	return false
}
