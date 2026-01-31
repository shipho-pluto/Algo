package main

import "fmt"

func main() {
	arr := []int{9, 6, 5, 1, 4, 2}
	target := 10
	fmt.Println(bufferSumK(arr, target))
}

func bufferSumK(arr []int, target int) (int, int) {
	n := len(arr)
	st := make(map[int]int)
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		if j, ok := st[sum-target]; ok {
			return j + 1, i
		}
		st[sum] = i
	}
	return -1, -1
}
