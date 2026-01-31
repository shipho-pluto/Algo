package main

import "fmt"

func main() {
	arr := []int{-10, -3, -1, 0, 3, 4, 5, 7, 10, 12}
	k, index := 4, 3
	fmt.Println(nearKToIndex(arr, k, index))
}

func nearKToIndex(arr []int, k int, index int) []int {
	n := len(arr)
	var res []int
	for i, l, r := k, index, index+1; i > 0; i-- {
		if l > 0 && r < n {
			if arr[index]-arr[l] < arr[r]-arr[index] {
				res = append(res, arr[l])
				l--
			} else {
				res = append(res, arr[r])
				r++
			}
		} else if l > 0 {
			res = append(res, arr[l])
			l--
		} else {
			res = append(res, arr[r])
			r++
		}
	}
	return res
}
