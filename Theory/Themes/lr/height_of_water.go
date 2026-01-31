package main

import "fmt"

func HeightOfWater(arr []int) int {
	sum := 0
	maxEl, maxIn := arr[0], 0
	n := len(arr)
	for i := range n {
		if arr[i] > maxEl {
			maxEl = arr[i]
			maxIn = i
		}
	}
	maxUp := arr[0]
	maxDown := arr[n-1]
	for i := 0; i < maxIn; i++ {
		if arr[i] >= maxUp {
			maxUp = arr[i]
		} else {
			sum += maxUp - arr[i]
		}

		if arr[n-1-i] >= maxDown {
			maxDown = arr[n-i-1]
		} else {
			sum += maxDown - arr[n-1-i]
		}
	}

	return sum
}

func main() {
	arr := []int{3, 1, 4, 3, 5, 1, 1, 3, 1}
	fmt.Println(HeightOfWater(arr))
}
