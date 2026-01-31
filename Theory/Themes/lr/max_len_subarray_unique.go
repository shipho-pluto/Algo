package main

import "fmt"

func maxLenUniqueElements(arr []int) int {
	n := len(arr)
	st := make(map[int]int)
	res := 0
	for l, r := 0, 0; l < n; l++ {
		r = max(r, st[arr[l]])
		res = max(res, l-r+1)
		st[arr[l]] = l + 1
	}

	return res
}

func main() {
	arr := []int{1, 2, 3, 1, 4, 5, 6}
	fmt.Println(maxLenUniqueElements(arr))
}
