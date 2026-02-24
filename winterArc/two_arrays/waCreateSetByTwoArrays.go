package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 2, 2, 4}
	arr2 := []int{2, 2, 4}

	fmt.Println(createSet(arr1, arr2))
	fmt.Println(createSetLR(arr1, arr2))
}

func createSet(arr1 []int, arr2 []int) []int {
	st := make(map[int]bool)
	for i := range arr1 {
		st[arr1[i]] = true
	}

	var res []int

	for i := range arr2 {
		if st[arr2[i]] {
			res = append(res, arr2[i])
			delete(st, arr2[i])
		}
	}

	return res
}

func createSetLR(arr1 []int, arr2 []int) []int {
	var res []int
	cur := arr1[0] - 1
	for l, r := 0, 0; l < len(arr1) && r < len(arr2); {
		if arr1[l] < arr2[r] {
			l++
		} else if arr1[l] > arr2[r] {
			r++
		} else if arr1[l] != cur {
			cur = arr1[l]
			res = append(res, arr1[l])
			l, r = l+1, r+1
		} else {
			l++
		}
	}
	return res
}
