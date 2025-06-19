package main

import "fmt"

func NumsEqZero(arr []int) int {
	pref := make(map[int]int)
	pref[0] = 1
	n := len(arr)
	sum := 0
	res := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
		if pref[sum] != 0 {
			res += pref[sum]
		}
		pref[sum]++
	}
	return res
}

func main() {
	arr := []int{1, -3, 2, 6, -5, 1, -2}
	fmt.Println(NumsEqZero(arr))
}
