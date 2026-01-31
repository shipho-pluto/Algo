package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3, 2, 0}
	arr2 := []int{5, 1, 2, 7, 3, 2}
	fmt.Println(multiplayArray(arr1, arr2))
}

func multiplayArray(arr1, arr2 []int) []int {
	n := len(arr1)
	st := make(map[int]int, n)
	for i := range n {
		st[arr1[i]]++
	}
	var res []int
	for i := range arr2 {
		if _, ok := st[arr2[i]]; ok {
			res = append(res, arr2[i])
			st[arr2[i]]--
			if st[arr2[i]] == 0 {
				delete(st, arr2[i])
			}
		}
	}
	return res
}
