package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, -1, -10, 4, 5, -1, 0}
	fmt.Println(findPrefSumX(arr, -3)) // max length
}

func findPrefSumX(arr []int, k int) [2]int {
	st := make(map[int]int)
	st[0] = 0

	inter := [2]int{-1, -1}
	sum := 0
	maxL := 0
	for i := range arr {
		sum += arr[i]
		if l, ok := st[sum-k]; ok {
			if i-l > maxL {
				maxL = i - l
				inter = [2]int{l, i}
			}
		}
		st[sum] = i + 1
	}
	return inter
}
