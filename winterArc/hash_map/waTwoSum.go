package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	target := 17
	fmt.Println(twoSum(arr, target))
}

func twoSum(arr []int, t int) []int {
	sum := make(map[int]int)
	for r := range arr {
		if l, ok := sum[t-arr[r]]; ok {
			return []int{l, r}
		}
		sum[arr[r]] = r
	}

	return []int{-1, -1}
}
