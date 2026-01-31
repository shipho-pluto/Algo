package main

import "fmt"

func main() {
	arr1 := []int{1, 2}
	arr2 := []int{1, 3, 5}
	fmt.Println(uniqElementsInFirst(arr1, arr2))
}

func uniqElementsInFirst(arr1 []int, arr2 []int) []int {
	var res []int

	for l, r := 0, 0; l < len(arr1); {
		if r < len(arr2) && arr1[l] < arr2[r] || r >= len(arr2) {
			res = append(res, arr1[l])
			l++
		} else if r < len(arr2) && arr1[l] > arr2[r] {
			r++
		} else {
			l++
			r++
		}
	}

	return res
}

//	st := make(map[int]bool)
//	for _, v := range arr2 {
//		st[v] = true
//	}
//	var res []int
//	for _, v := range arr1 {
//		if _, ok := st[v]; !ok {
//			res = append(res, v)
//		}
//	}
