package main

import "fmt"

func main() {
	arr := []int{-4, -3, -2, -1, 1, 2, 3, 12}
	target := 4
	fmt.Println(findTargetInSorted(arr, target))
}

func findTargetInSorted(arr []int, target int) bool {
	n := len(arr)
	sum := 0
	for l, r := 0, n-1; l < r; {
		sum = arr[l] + arr[r]
		if sum < target {
			l++
		} else if sum == target {
			fmt.Println(l, r)
			return true
		} else {
			r--
		}
	}
	return false
}
