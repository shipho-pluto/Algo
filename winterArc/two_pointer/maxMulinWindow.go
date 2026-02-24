package main

import "fmt"

func main() {
	arr := []int{1, 1, 3, 0, 3, 1, 0, 2, 1, 2, 4, 1}
	k := 3
	fmt.Println(maxMulInWindow(arr, k))
}

func maxMulInWindow(arr []int, k int) any {
	mul := 1
	maxMul := mul
	kc := k
	for i := range arr {
		if arr[i] == 0 {
			mul = 1
			kc = k
		} else {
			if kc == 0 {
				mul *= arr[i]
				mul /= arr[i-k]
				maxMul = max(maxMul, mul)
			} else {
				mul *= arr[i]
				kc--
			}
		}
	}

	return maxMul
}
