package main

import "fmt"

func main() {
	arr := []int{3, 1, 4, 3, 5, 1, 1, 3, 5}
	fmt.Println(highOfWater(arr))
}

func highOfWater(arr []int) int {
	maxHigh := 0
	imh := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] > maxHigh {
			maxHigh = arr[i]
			imh = i
		}
	}

	res := 0
	lm, rm := 0, 0

	for l, r := 0, n-1; l < imh || r > imh; l, r = l+1, r-1 {
		if l < imh {
			if arr[l] > lm {
				lm = arr[l]
			} else {
				res += lm - arr[l]
			}
		}

		if r > imh {
			if arr[r] > rm {
				rm = arr[r]
			} else {
				res += rm - arr[r]
			}
		}
	}

	return res
}
