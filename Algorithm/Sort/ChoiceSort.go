package main

import "fmt"

func ChoiceSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		_, id := MinElement(arr[i:])
		arr[i], arr[id+i] = arr[id+i], arr[i]
	}
	return arr
}

func MinElement(arr []int) (int, int) {
	minEl := arr[0]
	index := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < minEl {
			minEl = arr[i]
			index = i
		}
	}
	return minEl, index
}

func main() {
	fmt.Print(ChoiceSort([]int{123, 123, 12, 34453, 4567, 89, 56, 45, 345, 897, 67, 234}))
}
