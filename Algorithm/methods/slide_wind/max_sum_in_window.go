package main

import "fmt"

func main() {
	arr := []int{2, 9, -1, 0, 3, 4, -9, 0, 1, 2, -10, -11, 30}
	k := 4
	fmt.Println(maxSumWindow(arr, k))
}

func maxSumWindow(arr []int, k int) int {
	sum := 0
	n := len(arr)
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	res := sum
	for i := k; i < n; i++ {
		sum += arr[i] - arr[i-k]
		res = max(res, sum)
	}
	return res
}
