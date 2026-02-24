package main

import "fmt"

func main() {
	arr := []int{1, 1, 3, 0, 3, 1, 0, 1, 1}
	fmt.Println(maxMulInWindow(arr, 3))
	fmt.Println(maxMulInWindow2(arr, 3))
}

func maxMulInWindow(arr []int, k int) int {
	n := len(arr)
	mul, res := 1, 0
	delFlag := false
	for l, r := 0, 0; r < n; r++ {
		if arr[r] == 0 {
			mul = 1
			l = r + 1
			delFlag = false
		} else {
			mul *= arr[r]
			if r-l == k {
				delFlag = true
			}
			if delFlag {
				mul /= arr[r-k]
			}
			if r-l+1 == k {
				res = max(res, mul)
			}
		}
	}

	return res
}

func maxMulInWindow2(arr []int, k int) int {
	n := len(arr)
	mul, res, cnt := 1, 0, 0
	for i := 0; i < k; i++ {
		if arr[i] != 0 {
			mul *= arr[i]
		} else {
			cnt++
		}
	}
	if cnt == 0 {
		res = mul
	}
	for i := k; i < n; i++ {
		if arr[i] != 0 {
			mul *= arr[i]
		} else {
			cnt++
		}
		if arr[i-k] != 0 {
			mul /= arr[i-k]
		} else {
			cnt--
		}

		if cnt == 0 {
			res = max(res, mul)
		}
	}

	return res
}
