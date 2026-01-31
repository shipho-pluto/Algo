package main

import "fmt"

func main() {
	arr1 := []int{9, 3, 3, 1, 5, 4, 7, 2, 0}
	arr2 := []int{4, 5, 1, 2, 7, 5, 0, 9, 8}
	n := len(arr1)
	res := make([]int, n)
	c := 0

	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	for i := 0; i < n; i++ {
		if !set1[arr1[i]] {
			set1[arr1[i]] = true
			if set2[arr1[i]] {
				c++
			}
		}

		if !set2[arr2[i]] {
			set2[arr2[i]] = true
			if set1[arr2[i]] {
				c++
			}
		}

		res[i] = c
	}

	fmt.Println(res)
}
