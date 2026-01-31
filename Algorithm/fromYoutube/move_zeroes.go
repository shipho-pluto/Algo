package main

import "fmt"

func main() {
	arr := []int{1, 2, 0, 3, 0, 4, 0, 2, 0, 0, 1}
	fmt.Println(moveZeroes(arr))
}

func moveZeroes(arr []int) []int {
	n := len(arr)
	l := 0
	for ; l < n; l++ {
		if arr[l] == 0 {
			break
		}
	}
	if l == n {
		return arr
	}

	r := l + 1
	for ; l < n; l++ {
		for ; r < n; r++ {
			if arr[r] != 0 {
				break
			}
		}
		if r != n {
			arr[r], arr[l] = arr[l], arr[r]
		}
	}

	return arr
}
