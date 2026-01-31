package main

import "fmt"

func moveZeroesToEnd(arr []int) []int {
	n := len(arr)
	for r, l := 0, 0; l < n && r < n; {
		for ; r < n && arr[r] != 0; r++ {
		}
		l = r + 1
		for ; l < n && arr[l] == 0; l++ {
		}
		if l < n && r < n {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}

	return arr
}

func main() {
	arr := []int{1, 2, 3, 0, 1}
	fmt.Println(moveZeroesToEnd(arr))
}
