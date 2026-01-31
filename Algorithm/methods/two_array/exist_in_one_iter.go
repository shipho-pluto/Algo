package main

import "fmt"

func main() {
	arr1, arr2 := []int{1, 4, 6, 10, 0}, []int{4, 5, 6, 7, 10}
	fmt.Println(eioIter(arr1, arr2))
	fmt.Println(eioIterPre(arr1, arr2))
}

func eioIter(arr1 []int, arr2 []int) []int {
	var res, pre []int
	var ex [][]int
	c := 0
	st1, st2 := make(map[int]bool), make(map[int]bool)

	for i := range arr1 {
		st2[arr2[i]] = true

		if !st1[arr1[i]] {
			st1[arr1[i]] = true
			if !st2[arr1[i]] {
				c++
				pre = append(pre, arr1[i])
			}
		}

		res = append(res, c)
		ex = append(ex, pre)
	}

	fmt.Println(ex)
	return res
}

func eioIterPre(arr1 []int, arr2 []int) []int {
	var res, pre []int
	var ex [][]int
	c := 0
	st2 := make(map[int]bool)
	for i := range arr2 {
		st2[arr2[i]] = true
	}

	for i := range arr1 {
		if !st2[arr1[i]] {
			c++
			pre = append(pre, arr1[i])
		}

		res = append(res, c)
		ex = append(ex, pre)
	}

	fmt.Println(ex)
	return res
}
