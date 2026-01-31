package main

import "fmt"

func main() {
	arr := []int{1, -1, 3, 2, 3, -5, 0, -3}
	fmt.Println(cntSubSumZero(arr))
}

func cntSubSumZero(arr []int) int {
	n := len(arr)
	res := 0
	pref := make(map[int]int)
	pref[0] = 1
	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		if val, ok := pref[sum]; ok {
			res += val
		}
		pref[sum]++
	}
	return res
}
