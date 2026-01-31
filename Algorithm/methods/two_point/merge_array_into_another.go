package main

import "fmt"

func main() {
	arr := []int{2, 2, 4, 0, 0, 0, 0}
	arr2 := []int{1, 3, 6, 10}
	fmt.Println(mergeArrayInFirst(arr, arr2, 3, 4))
}

func mergeArrayInFirst(arr1 []int, arr2 []int, n, m int) []int {
	for i, l, r := m+n-1, n-1, m-1; l > -1 || r > -1; i-- {
		if l > -1 && r > -1 {
			if arr1[l] > arr2[r] {
				arr1[i] = arr1[l]
				l--
			} else {
				arr1[i] = arr2[r]
				r--
			}
		} else if l > -1 {
			arr1[i] = arr1[l]
			l--
		} else {
			arr1[i] = arr2[r]
			r--
		}
	}
	return arr1
}
