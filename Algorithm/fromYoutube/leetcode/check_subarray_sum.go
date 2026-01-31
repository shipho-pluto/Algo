package main

import "fmt"

func main() {
	arr := []int{23, 2, 6, 4, 7}
	k := 7
	fmt.Println(checkSubarraySum(arr, k))
}

func checkSubarraySum(arr []int, k int) bool {
	hm := make(map[int]int)
	hm[0] = -1
	sum := 0
	for i := range arr {
		sum += arr[i]
		mod := sum % k
		if _, ok := hm[mod]; ok {
			if i-hm[mod] >= 2 {
				return true
			}
		} else {
			hm[mod] = i
		}
	}
	return false
}
