package main

import "fmt"

func main() {
	arr1 := []int{0, 1, 6, -1, 9, 3, 2}
	arr2 := []int{9, 4, 5, 6, -1, 3, 0}
	fmt.Println(unionArray(arr1, arr2))
}

func unionArray(arr1 []int, arr2 []int) []int {
	n := len(arr1)
	res := make([]int, n)
	st1, st2 := make(map[int]bool), make(map[int]bool)
	c := 0

	for i := range arr2 {
		if !st1[arr1[i]] {
			st1[arr1[i]] = true
			if st2[arr1[i]] {
				c++
			}
		}

		if !st2[arr2[i]] {
			st2[arr2[i]] = true
			if st1[arr2[i]] {
				c++
			}
		}

		res[i] = c
	}

	return res
}
