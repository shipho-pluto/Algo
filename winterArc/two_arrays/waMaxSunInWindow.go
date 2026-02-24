package main

import "fmt"

func main() {
	arr := []int{2, 9, -1, 0, 3, 4, -9, 0, 1, 2, -10, -11, 30}
	k := 4
	fmt.Println(maxSumWindow(arr, k))
}

func maxSumWindow(arr []int, k int) any {
	sum := 0
	for i := range k {
		sum += arr[i]
	}
	maxSum := sum
	for i := k; i < len(arr); i++ {
		sum += arr[i] - arr[i-k]
		maxSum = max(maxSum, sum)
	}

	return maxSum
}
