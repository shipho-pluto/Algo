package main

import "fmt"

func main() {
	arr1, arr2 := []int{1, 4, 7, -1, 9, 0}, []int{5, 6, 1, 7, 8, 4}
	fmt.Println(esIter(arr1, arr2))
}

func esIter(arr1 []int, arr2 []int) []int {
	var res, pre []int
	var ex [][]int
	st1, st2 := make(map[int]bool), make(map[int]bool)
	c := 0
	for i := range arr1 {
		if !st1[arr1[i]] {
			st1[arr1[i]] = true
			if st2[arr1[i]] {
				c++
				pre = append(pre, arr1[i])
			}
		}

		if !st2[arr2[i]] {
			st2[arr2[i]] = true
			if st1[arr2[i]] {
				c++
				pre = append(pre, arr2[i])
			}
		}

		ex = append(ex, pre)
		res = append(res, c)
	}

	fmt.Println(ex)
	return res
}
