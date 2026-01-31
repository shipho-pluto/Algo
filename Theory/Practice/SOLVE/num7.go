package main

import "fmt"

func prefSum(arr []int, k int) (int, int) {
	n := len(arr)
	pref := make(map[int]int)
	pref[0] = -1
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		if data, ok := pref[sum-k]; ok {
			return data + 1, i
		}
		pref[sum] = i
	}
	return -1, -1
}

func main() {
	arr := []int{1, -11, 10, 4, 5, 9, 17, -20, 3, 18}
	k := 6
	fmt.Println(prefSum(arr, k))
}
