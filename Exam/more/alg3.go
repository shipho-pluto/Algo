package main

import "fmt"

func cntElementsInARowSumK(arr []int, k int) int {
	n := len(arr)
	cnt, sum := 0, 0
	r, l := 0, 0
	for ; l < n && r < n; r++ {
		if sum += arr[r]; sum == k {
			cnt++
			sum -= arr[l]
			l++
		} else if sum > k {
			sum -= arr[l]
			l++
		}
	}
	for ; l < n; l++ {
		if sum == k {
			cnt++
			sum -= arr[l]
			l++
		} else if sum > k {
			sum -= arr[l]
			l++
		}
	}
	return cnt
}

func main() {
	arr := []int{2, 2, 2, 2, 1}
	k := 2
	fmt.Println(cntElementsInARowSumK(arr, k))
}
