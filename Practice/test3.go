package main

import (
	"fmt"
)

func Abs3(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func findKClosest(arr []int, index, k int) []int {
	var result []int
	n := len(arr)
	if k < 1 || k > n || index > n-1 || index < 0 {
		return result
	}
	l := arr[index]
	r := l + 1
	for i := 0; i < k; i++ {
		if r == n || Abs3(arr[l]-arr[index]) < Abs3(arr[r]-arr[index]) {
			result = append(result, arr[l])
			l--
		} else if l < 0 || Abs3(arr[l]-arr[index]) >= Abs3(arr[r]-arr[index]) {
			result = append(result, arr[r])
			r++
		}
	}
	return result
}

func main() {
	a := []int{-4, -1, 3, 6, 9}
	index := 2
	k := 4

	result := findKClosest(a, index, k)
	fmt.Println(result)
}
