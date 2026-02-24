package main

import "fmt"

func main() {
	arr := []int{1, 2, 0, 3, 0, 4, 0, 2, 0, 0, 1}
	fmt.Println(moveZeros(arr))
}

func moveZeros(arr []int) []int {

	for l := 0; l < len(arr); {
		for ; arr[l] != 0; l++ {
		}
		r := l + 1
		for ; r < len(arr) && arr[r] == 0; r++ {
		}

		if r < len(arr) {
			arr[l], arr[r] = arr[r], arr[l]
		} else {
			break
		}
	}

	return arr
}
