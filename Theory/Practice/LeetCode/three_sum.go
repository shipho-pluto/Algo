package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-4, -3, 0, 1, 7, 11, 15, 21}
	target := 6
	fmt.Println(threeSum(arr, target))
}

func threeSum(arr []int, target int) []int {
	n := len(arr)
	sort.Ints(arr)
	minDif := arr[n-1]
	var res []int
	for i := 0; i < n-2; i++ {
		if i == 0 || i > 0 && arr[i] != arr[i-1] {
			for l, r := i+1, n-1; l < r; {
				if arr[r]+arr[l]+arr[i] > target {
					if minDif > (arr[r]+arr[l]+arr[i])-target {
						minDif = (arr[r] + arr[l] + arr[i]) - target
						res = []int{i, l, r, arr[r] + arr[l] + arr[i]}
					}
					r--
				} else if arr[r]+arr[l]+arr[i] < target {
					if minDif > target-(arr[r]+arr[l]+arr[i]) {
						minDif = target - (arr[r] + arr[l] + arr[i])
						res = []int{i, l, r, arr[r] + arr[l] + arr[i]}
					}
					l++
				} else {
					return []int{i, l, r, arr[r] + arr[l] + arr[i]}
				}
			}
		}
	}
	return res
}
