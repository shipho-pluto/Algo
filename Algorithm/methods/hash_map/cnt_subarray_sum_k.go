package main

import "fmt"

func main() {
	arr := []int{2, 2, 1, 1, 0, 2, -1}
	k := 2
	fmt.Println(cntSubarraySumK(arr, k))
}

func cntSubarraySumK(arr []int, k int) int {
	n := len(arr)

	sum := 0
	pref := make(map[int]int, n)
	pref[0] = 1
	res := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		if val, ok := pref[sum-k]; ok {
			res += val
		}
		pref[sum]++
	}
	return res
}
