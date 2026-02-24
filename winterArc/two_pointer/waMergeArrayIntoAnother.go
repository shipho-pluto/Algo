package main

import "fmt"

func main() {
	arr1 := []int{2, 2, 4, 0, 0, 0, 0}
	arr2 := []int{1, 3, 6, 10}
	mergeArrayIntoAnother(arr1, arr2, 3, 4)
	fmt.Println(arr1)
}

func mergeArrayIntoAnother(arr1 []int, arr2 []int, m int, n int) {
	l, r := m-1, n-1
	for i := m + n - 1; i > -1; i-- {
		if l > -1 && (r > -1 && arr1[l] >= arr2[r] || r == -1) {
			arr1[i] = arr1[l]
			l--
		} else {
			arr1[i] = arr2[r]
			r--
		}
	}
}
