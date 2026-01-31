package main

import "fmt"

func main() {
	arr := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(arr))
}

func trap(arr []int) int {
	var res int
	maxHigh := 0
	maxIndex := 0
	for i := range arr {
		if maxHigh < arr[i] {
			maxHigh = arr[i]
			maxIndex = i
		}
	}
	maxBefore, maxAfter := arr[0], arr[len(arr)-1]
	for l, r := 0, len(arr)-1; l < maxIndex || r > maxIndex; {
		maxBefore = max(maxBefore, arr[l])
		maxAfter = max(maxAfter, arr[r])
		if l < maxIndex {
			res += maxBefore - arr[l]
			l++
		}
		if r > maxIndex {
			res += maxAfter - arr[r]
			r--
		}
	}
	return res
}
