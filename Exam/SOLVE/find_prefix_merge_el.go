package main

import "fmt"

func findPrefixMergeElements(arr1, arr2 []int) []int {
	n := len(arr1)
	s1, s2 := map[int]bool{}, map[int]bool{}

	c := 0
	res := make([]int, n)
	for i := 0; i < n; i++ {
		if !s1[arr1[i]] {
			s1[arr1[i]] = true
			if s2[arr1[i]] == true {
				c++
			}
		}

		if !s2[arr2[i]] {
			s2[arr2[i]] = true
			if s1[arr2[i]] {
				c++
			}
		}
		res[i] = c
	}
	return res
}

func main() {
	s1, s2 := []int{1, 3, 5, 1, 9, 9, 0, 0, 2}, []int{3, 5, 2, 2, 9, 1, 2, 2, 0}
	fmt.Println(findPrefixMergeElements(s1, s2))
}
