package main

import "fmt"

func main() {
	arr := []int{0, 0, 1, 0, 2, 0, 3}
	fmt.Println(replaceZeroes(arr))
}

func replaceZeroes(arr []int) []int {
	n := len(arr)
	for l := 0; l < n; l++ {
		for ; l < n; l++ {
			if arr[l] == 0 {
				break
			}
		}
		r := l + 1
		for ; r < n; r++ {
			if arr[r] != 0 {
				break
			}
		}
		if l < n && r < n {
			arr[r], arr[l] = arr[l], arr[r]
		}
	}
	return arr
}
