package main

import "fmt"

func MergeSort(arr []int) []int {
	currentScr := arr
	currentDest := make([]int, len(arr))

	size := 1
	for size < len(arr) {
		for i := 0; i < len(arr); i += 2 * size {
			Merge(currentScr, currentScr, currentDest, i, i+size, i, size)
		}
		currentDest, currentScr = currentScr, currentDest
		size *= 2
	}
	return currentScr
}

func Merge(arr1, arr2, dest []int, st1, st2, dst, size int) {
	in1 := st1
	in2 := st2

	ed1 := min(st1+size, len(arr1))
	ed2 := min(st2+size, len(arr2))

	iteration := ed1 - in1 + ed2 - in2

	for i := dst; i < iteration+dst; i++ {
		if in1 < ed1 && (in2 == ed2 || arr1[in1] < arr2[in2]) {
			dest[i] = arr1[in1]
			in1++
		} else {
			dest[i] = arr2[in2]
			in2++
		}
	}
}

func main() {
	fmt.Print(MergeSort([]int{123, 123, 12, 34453, 4567, 89, 56, 45, 345, 897, 67, 234, 670}))
}
