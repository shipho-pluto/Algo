package main

import "fmt"

func main() {
	arr1, arr2 := []int{1, 3, 5, 0, 0}, []int{1, 2}
	fmt.Println(maio(arr1, arr2, 3, 2))
}

func maio(arr1 []int, arr2 []int, l1, l2 int) []int {
	for i, l, r := l1+l2-1, l1-1, l2-1; l >= 0 || r >= 0; i-- {
		if l >= 0 && r >= 0 {
			if arr1[l] > arr2[r] {
				arr1[i] = arr1[l]
				l--
			} else {
				arr1[i] = arr2[r]
				r--
			}
		} else if l >= 0 {
			arr1[i] = arr1[l]
			l--
		} else {
			arr1[i] = arr2[r]
			r--
		}
	}
	return arr1
}
