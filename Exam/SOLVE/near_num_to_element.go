package main

import "fmt"

func nearNumToElement(arr []int) int {
	n := len(arr)
	i := 0
	for ; i < n; i++ {
		if arr[i] == 0 {
			break
		}
	}

	res := i

	for j := i + 1; j < n; j++ {
		flag := false
		for ; j < n; j++ {
			if arr[j] == 1 {
				flag = true
			}
			if arr[j] == 0 {
				break
			}
		}
		for k := i + 1; k < j; k++ {
			if arr[k] == 1 {
				res = max(res, min(j-k, k-i))
			}
		}

		if j == n && flag {
			res = max(res, j-i-1)
		}
		i = j
	}

	return res
}

func main() {
	arr := []int{0, 1, 1, 1, 1, 1, 0, 0, 1, 2, 1, 1, 0, 2, 2, 2, 2, 2, 2}
	fmt.Println(nearNumToElement(arr))
}
