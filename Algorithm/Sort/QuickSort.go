package main

import "fmt"

func QuickSort(arr []int, left, right int) {
	if left < right {
		start := Part(arr, left, right)
		QuickSort(arr, left, start-1)
		QuickSort(arr, start, right)
	}
}

func Part(arr []int, left, right int) int {
	r := right
	l := left

	// меняем местами элемент больше опорного на элемент меньше опорного, чтобы получить 2 части
	pivot := arr[l]
	for l <= r {
		for ; arr[l] < pivot; l++ {
		}
		for ; arr[r] > pivot; r-- {
		}

		if l <= r {
			arr[l], arr[r] = arr[r], arr[l]
			l++
			r--
		}
	}
	return l
}

func main() {
	arr := []int{123, 123, 12, 34453, 4567, 89, 56, 45, 345, 897, 67, 234}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
