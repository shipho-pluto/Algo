package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 4, 6, 6, 8, 91}
	arr2 := []int{1, 5, 6, 10}

	fmt.Println(ex1FromAnother(arr1, arr2))
	fmt.Println(ex1FromAnotherCount(arr1, arr2))
	fmt.Println(ex1FromAnotherLR(arr1, arr2))
	fmt.Println(ex1FromAnotherLRCount(arr1, arr2))
}

func ex1FromAnother(arr1 []int, arr2 []int) any {
	st := make(map[int]int)

	for i := range arr2 {
		st[arr2[i]] = 1
	}

	var res []int
	for i := range arr1 {
		if _, ok := st[arr1[i]]; !ok {
			res = append(res, arr1[i])
		}
	}
	return res
}

func ex1FromAnotherCount(arr1 []int, arr2 []int) any {
	st := make(map[int]int)

	for i := range arr2 {
		st[arr2[i]]++
	}

	var res []int
	for i := range arr1 {
		if _, ok := st[arr1[i]]; ok {
			st[arr1[i]]--
			if st[arr1[i]] == 0 {
				delete(st, arr1[i])
			}
		} else {
			res = append(res, arr1[i])
		}
	}
	return res
}

func ex1FromAnotherLR(arr1 []int, arr2 []int) []int {
	var res []int
	for l, r := 0, 0; l < len(arr1); {
		if l < len(arr1) && r < len(arr2) {
			if arr1[l] < arr2[r] {
				res = append(res, arr1[l])
				l++
			} else if arr1[l] > arr2[r] {
				r++
			} else {
				l++
			}
		} else {
			res = append(res, arr1[l])
			l++
		}
	}

	return res
}

func ex1FromAnotherLRCount(arr1 []int, arr2 []int) []int {
	var res []int
	for l, r := 0, 0; l < len(arr1); {
		if l < len(arr1) && r < len(arr2) {
			if arr1[l] < arr2[r] {
				res = append(res, arr1[l])
				l++
			} else if arr1[l] > arr2[r] {
				r++
			} else {
				l, r = l+1, r+1
			}
		} else {
			res = append(res, arr1[l])
			l++
		}
	}

	return res
}
