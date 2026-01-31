package main

import "fmt"

func main() {
	arr := []int{1, 0, 0, 0, 4, 5, 6, 0, 0, 1, 0, 2}
	moveZeroes(arr)
	fmt.Println(arr)
}

func moveZeroes(arr []int) {
	n := len(arr)
	for l := 0; l < n; l++ {
		for l < n && arr[l] != 0 {
			l++
		}
		r := l + 1
		for r < n && arr[r] == 0 {
			r++
		}
		if l != n && r != n {
			arr[l], arr[r] = arr[r], arr[l]
		}
	}
}
