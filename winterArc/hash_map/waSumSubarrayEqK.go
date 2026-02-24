package main

import "fmt"

func main() {
	arr := []int{9, 6, 5, 1, 4, 2}
	fmt.Println(findTarget(arr, 10))
}

func findTarget(arr []int, k int) (int, int) {
	buf := make(map[int]int)
	buf[0] = 0
	sum := 0
	for r := range arr {
		sum += arr[r]
		if l, ok := buf[sum-k]; ok {
			return l, r
		}

		buf[sum] = r + 1
	}
	return -1, -1
}
