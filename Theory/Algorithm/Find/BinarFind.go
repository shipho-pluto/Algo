package main

import "fmt"

// Любая сортировка

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

func BinaryFindReq(arr []int, start, end, target int) int {
	middleIndex := start + (end-start)/2
	average := arr[middleIndex]
	if average == target {
		return middleIndex
	}
	if average > target {
		return BinaryFindReq(arr, start, middleIndex-1, target)
	} else if average < target {
		return BinaryFindReq(arr, middleIndex+1, end, target)
	}
	return -1

}

func BinaryFindFor(arr []int, target int) int {
	startIndex := 0
	endIndex := len(arr) - 1
	var middleIndex int
	for startIndex <= endIndex {
		middleIndex = (endIndex - startIndex) / 2

		if arr[middleIndex] == target {
			return middleIndex
		}

		if arr[middleIndex] > target {
			endIndex = middleIndex - 1
		} else {
			startIndex = middleIndex + 1
		}
	}
	return -1
}

func main() {
	arr := []int{123, 123, 12, 34453, 4567, 89, 56, 45, 345, 897, 67, 234}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Print(BinaryFindReq(arr, 0, len(arr), 12))
}
