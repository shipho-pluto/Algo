package main

import "fmt"

func maxSlidingWindow(arr []int, k int) []int {
	n := len(arr)
	maxEl := arr[0]
	maxIn := 0
	for i := 0; i < k; i++ {
		if arr[i] > maxEl {
			maxEl = arr[i]
			maxIn = i
		}
	}
	res := make([]int, n-k+1)
	res[0] = maxEl
	for i := 1; i <= n-k; i++ {
		window := arr[i : i+k]
		if window[k-1] >= maxEl {
			maxEl = window[k-1]
			maxIn = i + k - 1
		} else if maxIn == i-1 {
			maxEl = window[0]
			maxIn = i
			for j, w := range window {
				if w > maxEl {
					maxEl = w
					maxIn = j + i
				}
			}
		}
		res[i] = maxEl
	}
	return res
}

func main() {
	arr := []int{-6, -10, -7, -1, -9, 9, -8, -4, 10, -5, 2, 9, 0, -7, 7, 4, -2, -10, 8, 7}
	k := 7
	fmt.Println(maxSlidingWindow(arr, k))
}
