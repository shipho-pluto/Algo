package main

import "fmt"

func findTarget(arr []int, target int) (int, int) {
	n := len(arr)
	pref := make(map[int]int, n)
	pref[0] = -1
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		if val, ok := pref[sum-target]; ok {
			return val + 1, i
		}
		pref[sum] = i
	}

	return -1, -1
}

func main() {
	arr := []int{10, 6, 5, 1, 4, 2}
	fmt.Println(findTarget(arr, 10))
}
