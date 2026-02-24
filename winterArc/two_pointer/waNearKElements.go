package main

import "fmt"

func main() {
	arr := []int{-1, 0, 2, 3, 5, 6, 9}
	index, k := 2, 5

	fmt.Println(nearKElements(arr, index, k))
}

func abc(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func nearKElements(arr []int, index int, k int) []int {
	res := make([]int, k)
	for c, l, r := 0, index, index+1; c < k; c++ {
		if l > -1 && (r < len(arr) && abc(arr[r]-arr[index]) > abc(arr[l]-arr[index]) || r == len(arr)) {
			res[c] = arr[l]
			l--
		} else {
			res[c] = arr[r]
			r++
		}
	}
	return res
}
