package main

import "fmt"

func main() {
	arr := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(sorted2Sum(arr, target))
}

func sorted2Sum(arr []int, target int) []int {
	for l, r := 0, len(arr)-1; l < r; {
		if arr[l]+arr[r] < target {
			l++
		} else if arr[l]+arr[r] > target {
			r--
		} else {
			return []int{l + 1, r + 1}
		}
	}
	return []int{}
}
