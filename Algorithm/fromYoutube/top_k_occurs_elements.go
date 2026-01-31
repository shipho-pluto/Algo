package main

import "fmt"

func main() {
	arr := []int{1, 4, 5, 0, 2, 1, 1, 4, 5, 3, 9, 0, 0}
	fmt.Println(topKOccursElements(arr, 3))
}

func topKOccursElements(arr []int, k int) []int {
	res := make([]int, k)
	st := make(map[int]int)
	maxOccur := 0
	for i := range arr {
		st[arr[i]]++
		maxOccur = max(maxOccur, st[arr[i]])
	}
	occur := make([][]int, maxOccur)
	for key, val := range st {
		occur[val-1] = append(occur[val-1], key)
	}
	for i := maxOccur - 1; i >= 0; i-- {
		for j := range occur[i] {
			if k > 0 {
				res[k-1] = occur[i][j]
				k--
			}
		}
	}
	return res
}
