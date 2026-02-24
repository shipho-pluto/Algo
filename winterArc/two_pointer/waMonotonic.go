package main

import "fmt"

func main() {
	arr := []int{0, 1, 2, 3, 4, 10, 0, 2, 11, 1, -1, -2, -10, -11, -12}
	fmt.Println(maxLenMonotonicList(arr))
}

func maxLenMonotonicList(arr []int) (int, int) {
	n := len(arr)
	l, r := 0, 1
	increase := arr[l] < arr[r]
	maxLength := []int{0, 0, 0}
	for r < n {
		maxLength[1] = r
		maxLength[2] = max(maxLength[2], r-l)

		if increase && arr[r] >= arr[r-1] {
			r++
		} else if increase && arr[r] < arr[r-1] {
			increase = false
			l = r - 1
		}

		if !increase && arr[r] <= arr[r-1] {
			r++
		} else if !increase && arr[r] > arr[r-1] {
			increase = true
			l = r - 1
		}

		maxLength[0] = l
	}

	return maxLength[0], maxLength[1]
}
