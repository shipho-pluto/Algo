package main

import "fmt"

func main() {
	arr1, arr2 := []int{1, 4, 7, -1, 9, 0, 1}, []int{5, 6, 1, 1, 7, 8, 4}
	fmt.Println(equalElementsIter(arr1, arr2))
	fmt.Println(equalElementsIterCount(arr1, arr2))
}

func equalElementsIter(arr1 []int, arr2 []int) []int {
	st1, st2 := make(map[int]bool), make(map[int]bool)
	res := make([]int, len(arr1))
	var nums [][]int
	var p []int
	c := 0
	for i := range arr1 {
		if !st1[arr1[i]] {
			st1[arr1[i]] = true
			if st2[arr1[i]] {
				p = append(p, arr1[i])
				c++
			}
		}

		if !st2[arr2[i]] {
			st2[arr2[i]] = true
			if st1[arr2[i]] {
				p = append(p, arr2[i])
				c++
			}
		}

		nums = append(nums, p)
		res[i] = c
	}

	fmt.Println(nums)
	return res
}

func equalElementsIterCount(arr1 []int, arr2 []int) []int {
	st1, st2 := make(map[int]int), make(map[int]int)
	res := make([]int, len(arr1))
	var nums [][]int
	var p []int
	c := 0
	for i := range arr1 {
		st1[arr1[i]]++
		if st2[arr1[i]] > 0 {
			p = append(p, arr1[i])
			st2[arr1[i]]--
			c++
		}

		st2[arr2[i]]++
		if st1[arr2[i]] > 0 {
			p = append(p, arr2[i])
			st1[arr2[i]]--
			c++
		}

		nums = append(nums, p)
		res[i] = c
	}

	fmt.Println(nums)
	return res
}
