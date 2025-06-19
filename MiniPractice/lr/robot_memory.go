package main

import "fmt"

func RobotsMemory(arr string, k int) int {
	res := 0
	cnt := 0
	for i := k; i < len(arr); i++ {
		if arr[i] == arr[i-k] {
			cnt++
			res += cnt
		} else {
			cnt = 0
		}
	}
	return res
}

func main() {
	n := "abcabcacc"
	k := 3
	fmt.Println(RobotsMemory(n, k))
}
