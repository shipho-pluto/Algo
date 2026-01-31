package main

import (
	"fmt"
)

func NearShopToHouse(arr []int) int {
	n := len(arr)
	zero := 0

	for i := 0; i < n && arr[i] != 1; i++ {
		if arr[i] == 0 {
			zero = i
		}
	}

	res := 0
	for i := zero + 1; i < n; i++ {
		iEnd := 0
		for i < n && arr[i] != 0 {
			if arr[i] == 1 {
				iEnd = i
			}
			i++
		}
		if arr[zero] != 0 {
			res = i
		}
		if i == n {
			res = max(iEnd-zero, res)
		}
		for j := zero + 1; j < i; j++ {
			c := 0
			if arr[j] == 1 {
				c = min(j-zero, i-j)
				res = max(res, c)
			}
		}
		zero = i
	}

	return res
}

func main() {
	arr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 0, 1, 2, 1, 1, 0, 1}
	fmt.Println(NearShopToHouse(arr))
}
