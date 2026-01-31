package main

import "fmt"

func main() {
	arr := []int{4, 3, 1}
	fmt.Println(checkPossibility(arr))
}

func checkPossibility(arr []int) bool {
	count := 0
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			if i > 0 && arr[i-1] > arr[i+1] && i < len(arr)-2 && arr[i] > arr[i+1] {
				return false
			}
			count++
		}
	}
	return count <= 1
}
