package main

import "fmt"

func main() {
	arr := []int{1, 3, 4, -1, -10, 4, 5, -1}
	fmt.Println(findPrefSumX(arr, -3))
}

func findPrefSumX(arr []int, k int) (int, int) {
	n := len(arr)
	pref := make(map[int]int)
	pref[0] = -1
	sum := 0
	for i := range n {
		sum += arr[i]
		if val, ok := pref[sum-k]; ok {
			return val + 1, i
		}
		pref[sum] = i
	}

	return -1, -1
}
