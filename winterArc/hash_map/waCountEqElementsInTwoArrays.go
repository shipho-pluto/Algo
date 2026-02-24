package main

import "fmt"

func main() {
	arr1 := []int{1, 1, 2, 3, 4}
	arr2 := []int{2, 1, 0, 0, 1}
	fmt.Println(cntEqEl(arr1, arr2))
}

func cntEqEl(arr1 []int, arr2 []int) []int {
	res := make([]int, len(arr1))
	st1, st2 := make(map[int]bool), make(map[int]bool)
	c := 0
	for i := range arr1 {
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
