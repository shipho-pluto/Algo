package main

import "fmt"

func main() {
	arr := []int{-5, -2, -1, 0, 1, 2, 5, 7, 19}
	k := 100
	fmt.Println(twoSum(arr, k))
}

func twoSum(arr []int, k int) (int, int) {
	n := len(arr)
	for l, r := 0, n-1; l < r; {
		if arr[l]+arr[r] < k {
			l++
		} else if arr[l]+arr[r] > k {
			r--
		} else {
			return l, r
		}
	}
	return -1, -1
}
