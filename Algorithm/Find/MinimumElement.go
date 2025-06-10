package main

import "fmt"

func MinElement(array []int) (int, int) {

	minElem := array[0]
	index := 0
	for i := 1; i < len(array); i++ {
		if array[i] < minElem {
			minElem = array[i]
			index = i
		}
	}
	return array[index], index
}

func main() {
	fmt.Print(MinElement([]int{123, 123, 12, 34453, 4567, 89, 56, 45, 345, 897, 67, 234}))
}
