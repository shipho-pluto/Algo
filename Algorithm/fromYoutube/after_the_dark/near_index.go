package main

import "fmt"

func main() {
	arr := []int{-10, -9, -3, -1, 0, 2, 5, 8, 19}
	k, index := 4, 3
	fmt.Println(nearKToIndex(arr, k, index))
}

func nearKToIndex(arr []int, k int, index int) []int {
	l, r := index, index+1
	res := make([]int, k)
	for i := 0; i < k; i++ {
		if l > -1 && r < len(arr)-1 {
			if arr[index]-arr[l] < arr[r]-arr[index] {
				res[i] = arr[l]
				l--
			} else {
				res[i] = arr[r]
				r++
			}
		} else if l > -1 {
			res[i] = arr[l]
			l--
		} else {
			res[i] = arr[r]
			r++
		}
	}
	return res
}
