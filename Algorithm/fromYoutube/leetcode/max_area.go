package main

import "fmt"

func main() {
	arr := []int{8, 7, 2, 1}
	fmt.Println(maxArea(arr))
}

func maxArea(arr []int) int {
	res := 0
	for l, r := 0, len(arr)-1; l < r; {
		res = max(res, (r-l)*(min(arr[l], arr[r])))
		if arr[l] < arr[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
