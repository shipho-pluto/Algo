package main

import "fmt"

func cntBufferSumK(arr []int, k int) int {
	n := len(arr)
	sum, cnt := 0, 0
	pref := make(map[int]int)
	pref[0] = -1
	for i := 0; i < n; i++ {
		sum += arr[i]
		if _, ok := pref[sum-k]; ok {
			cnt++
		}
		pref[sum] = i
	}

	return cnt
}

func main() {
	arr := []int{1, 4, 2, 9, 1, 8, -10, 9, -1, -2, 9}
	k := 17
	fmt.Println(cntBufferSumK(arr, k))
}
