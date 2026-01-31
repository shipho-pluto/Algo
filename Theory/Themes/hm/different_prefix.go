package main

import "fmt"

func Add1DifferentPrefix(arr []int, l, r []int) []int {
	n := len(arr)
	mask := make([]int, n)
	for i := 0; i < len(l); i++ {
		mask[l[i]]++
		if r[i]+1 != n {
			mask[r[i]+1]--
		}

	}
	res := make([]int, n)
	cnt := 0
	for i := 0; i < n; i++ {
		cnt += mask[i]
		res[i] = arr[i] + cnt
	}

	return res
}

func main() {
	arr := []int{5, 1, 9, 4, 0, 2, 4, -1, 8}
	l, r := []int{3}, []int{7}
	fmt.Println(Add1DifferentPrefix(arr, l, r))
}
