package main

import (
	st "Algoritm/Algorithm/methods/struct/Structs"
	"fmt"
)

func main() {
	arr := []int{-1, 3, 2, 9, 4, 5, 3}
	k := 3
	fmt.Println(maxSlidingWindow(arr, k))
}

func maxSlidingWindow(arr []int, k int) []int {
	n := len(arr)
	res := make([]int, n-k+1)

	dq := st.NewDeQueue(true)
	dq.PutBack(arr[0])
	for i := 1; i < n; i++ {
		if i >= k {
			res[i-k] = dq.TopFront()
			if arr[i-k] == dq.TopFront() {
				dq.PopFront()
			}
		}
		for arr[i] >= dq.TopBack() {
			dq.PopBack()
		}
		dq.PutBack(arr[i])
	}
	res[n-k] = dq.TopFront()

	return res
}
