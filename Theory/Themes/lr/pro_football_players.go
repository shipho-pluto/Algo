package main

import "fmt"

func ProFootballPlayers(arr []int) int {
	pref := make([]int, len(arr)+1)
	n := len(arr)
	for i := 0; i < n; i++ {
		pref[i+1] = arr[i] + pref[i]
	}
	l, r := n-3, n-1
	ms := 0
	for ; r > 1; r-- {
		for ; l > -1 && arr[l]+arr[l+1] >= arr[r]; l-- {
		}
		ms = max(ms, pref[r+1]-pref[l+1])
	}
	return ms
}

func main() {
	arr := []int{3, 4, 4, 4, 4, 5, 5, 5, 6, 6, 7, 8, 12}
	fmt.Println(ProFootballPlayers(arr))
}
