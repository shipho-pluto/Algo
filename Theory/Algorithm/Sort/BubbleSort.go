package main

import "fmt"

func BubbleSort(arr []int) []int {

	for cp := 1; cp != 0; {
		cp = 0
		for i := 0; i < len(arr)-2; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				cp = 1
			}
		}
	}

	return arr
}

func main() {
	fmt.Print(BubbleSort([]int{123, 123, 12, 34453, 4567, 89, 56, 45, 345, 897, 67, 234}))
}
