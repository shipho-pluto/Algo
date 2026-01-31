package main

import "fmt"

func sortArrs(arr1, arr2 []int, m, n int) {
	l, r := m-1, n-1
	for i := m + n - 1; l > -1 || r > -1; i-- {
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
}

func main() {
	arr1, arr2 := []int{1, 4, 9, 10, 11, 0, 0, 0}, []int{-10, -2, 4}
	m, n := 5, 3
	sortArrs(arr1, arr2, m, n)
	fmt.Println(arr1)
}
