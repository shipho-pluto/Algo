package main

import "fmt"

func main() {
	arr1 := []int{1, 2, 3}
	arr2 := []int{2, 5, 6, 7, 8}
	fmt.Println(zigzagIt(arr1, arr2))
}

func zigzagIt(arr1 []int, arr2 []int) []int {
	res := make([]int, len(arr1)+len(arr2))
	for i, l, r := 0, 0, 0; i < len(res); i++ {
		if l < len(arr1) {
			res[i] = arr1[l]
			i++
			l++
		}
		if r < len(arr2) {
			res[i] = arr2[r]
			r++
		}
	}

	return res
}
