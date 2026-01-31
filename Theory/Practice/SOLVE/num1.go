package main

import "fmt"

func ShowNearKNumbers(arr []int, index, k int) []int {
	var res []int
	n := len(arr)
	cnt := 0
	l, r := index, index+1
	for l > -1 || r < n {
		if l > -1 && (r == n || arr[index]-arr[l] <= arr[r]-arr[index]) {
			res = append(res, arr[l])
			l--
			cnt++
		} else if r < n && (l == -1 || arr[index]-arr[l] > arr[r]-arr[index]) {
			res = append(res, arr[r])
			r++
			cnt++
		}

		if cnt == k {
			break
		}
	}
	return res
}

func main() {
	a := []int{-1, 3, 5, 10, 11, 19}
	k := 5
	index := 2
	fmt.Println(ShowNearKNumbers(a, index, k))
}
