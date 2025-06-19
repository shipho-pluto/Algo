package main

import "fmt"

func findPassNumber(arr []int) int {
	n := len(arr)
	sum := 0
	for i := range n {
		sum += arr[i]
	}
	return (n+1)*(n+2)/2 - sum
}

func findPassNumber2(arr []int, k int) []int {
	n := len(arr)
	sum := make([]int, n+k)
	var res []int
	for i := range n {
		sum[arr[i]-1]++
	}
	for i := range sum {
		if sum[i] == 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	arr := []int{9, 8, 4, 1, 7, 2, 5, 3}
	fmt.Println(findPassNumber(arr))
	arr = []int{9, 8, 4, 2, 5, 3}
	k := 3
	fmt.Println(findPassNumber2(arr, k))
}
