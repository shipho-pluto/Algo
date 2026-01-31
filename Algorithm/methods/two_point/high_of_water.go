package main

import "fmt"

func main() {
	arr := []int{3, 1, 4, 3, 5, 1, 1, 3, 5}
	fmt.Println(howMuchWater(arr))
}

func howMuchWater(arr []int) int {
	n := len(arr)
	maxHigh := arr[0]
	index := 0
	for i := 0; i < n; i++ {
		if maxHigh < arr[i] {
			maxHigh = arr[i]
			index = i
		}
	}
	maxBefore := arr[0]
	maxAfter := arr[n-1]
	res := 0
	for l, r := 0, n-1; l < index || r > index; l, r = l+1, r-1 {
		if l < index {
			maxBefore = max(maxBefore, arr[l])
			if arr[l] < maxBefore {
				res += maxBefore - arr[l]
			}
		}
		if r > index {
			maxAfter = max(maxAfter, arr[r])
			if arr[r] < maxAfter {
				res += maxAfter - arr[r]
			}
		}
	}

	return res
}
