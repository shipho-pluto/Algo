package main

import "fmt"

func main() {
	arr := []int{-10, -3, -1, 0, 2, 5, 19, 21}
	k := 6
	fmt.Println(sumK(arr, k))
}

func sumK(arr []int, k int) (int, int) {
	n := len(arr)
	hm := make(map[int]int)
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		if j, ok := hm[sum-k]; ok {
			return j + 1, i
		}
		hm[sum] = i
	}
	return -1, -1
}
