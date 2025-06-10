package main

import "fmt"

func LinearFind(arr []int, target int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			return i
		}
	}
	return -1
}

func main() {
	arr := []int{123, 123, 12, 34453, 4567, 89, 56, 45, 345, 897, 67, 234}
	fmt.Print(LinearFind(arr, 12))
}
