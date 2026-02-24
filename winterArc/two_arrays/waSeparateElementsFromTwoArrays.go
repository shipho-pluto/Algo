package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 4, 6, 6, 8, 91}
	arr2 := []int{1, 5, 6, 10}
	fmt.Println(separateArrays(arr1, arr2))
	fmt.Println(separateArraysLR(arr1, arr2))
	fmt.Println(separateArraysLRCount(arr1, arr2))

}

func separateArraysLR(arr1 []int, arr2 []int) []int {
	var res []int
	cur := arr1[0] - 1
	for l, r := 0, 0; l < len(arr1) || r < len(arr2); {
		if l < len(arr1) && r < len(arr2) {
			if arr1[l] < arr2[r] && arr1[l] != cur {
				res = append(res, arr1[l])
				l++
			} else if arr1[l] > arr2[r] && arr2[r] != cur {
				res = append(res, arr2[r])
				r++
			} else if arr1[l] == arr2[r] {
				cur = arr1[l]
				l, r = l+1, r+1
			} else if arr1[l] == cur {
				l++
			} else {
				r++
			}
		} else if l < len(arr1) {
			res = append(res, arr1[l])
			l++
		} else {
			res = append(res, arr2[r])
			r++
		}
	}

	return res
}

func separateArraysLRCount(arr1 []int, arr2 []int) []int {
	var res []int
	for l, r := 0, 0; l < len(arr1) || r < len(arr2); {
		if l < len(arr1) && r < len(arr2) {
			if arr1[l] < arr2[r] {
				res = append(res, arr1[l])
				l++
			} else if arr1[l] > arr2[r] {
				res = append(res, arr2[r])
				r++
			} else {
				l, r = l+1, r+1
			}
		} else if l < len(arr1) {
			res = append(res, arr1[l])
			l++
		} else {
			res = append(res, arr2[r])
			r++
		}
	}
	return res
}

func separateArrays(arr1 []int, arr2 []int) []int {
	st1, st2 := make(map[int]int), make(map[int]int)

	var res []int

	for i := range arr1 {
		st1[arr1[i]] = 1
	}

	for i := range arr2 {
		if _, ok := st1[arr2[i]]; !ok {
			res = append(res, arr2[i])
		}
		st2[arr2[i]] = 1
	}

	for i := range arr1 {
		if _, ok := st2[arr1[i]]; !ok {
			res = append(res, arr1[i])
		}
	}

	return res
}
