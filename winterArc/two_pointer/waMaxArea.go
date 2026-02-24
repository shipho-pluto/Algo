package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, 5, 2, 3, 9, 9, 1, 2, 3, 4, 8}
	fmt.Println(maxArea(arr))
}

func maxArea(arr []int) any {
	res := 0

	for l, r := 0, len(arr)-1; l < r; {
		res = max(res, (r-l)*min(arr[l], arr[r]))
		if arr[l] < arr[r] {
			l++
		} else {
			r--
		}
	}

	return res
}
