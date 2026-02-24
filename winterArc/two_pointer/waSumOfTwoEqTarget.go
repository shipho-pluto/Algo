package main

import "fmt"

func main() {
	arr := []int{-10, -3, -1, 0, 1, 2, 3, 6, 8, 11}
	target := 10
	fmt.Println(sumOfTwoEqTarget(arr, target))
}

func sumOfTwoEqTarget(arr []int, target int) bool {
	for l, r := 0, len(arr)-1; l < r; {
		if arr[l]+arr[r] < target {
			l++
		} else if arr[l]+arr[r] > target {
			r--
		} else {
			fmt.Println(l, r)
			return true
		}
	}

	return false
}
