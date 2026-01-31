package main

import "fmt"

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)
}

func sortColors(arr []int) {
	for i, l, r := 0, 0, len(arr)-1; i <= r; i++ {
		if arr[i] == 0 {
			arr[i], arr[l] = arr[l], arr[i]
			l++
		} else if arr[i] == 2 {
			arr[i], arr[r] = arr[r], arr[i]
			i--
			r--
		}
	}
}
