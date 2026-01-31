package main

import (
	"Algoritm/Theory/Sructure"
	"fmt"
)

func maxSlidingWindow(arr []int, k int) []int {
	if k == 1 {
		return arr
	}
	n := len(arr)
	dqueue := Struct.DQueue{}
	dqueue.PushFront(arr[0])
	for i := 1; i < k; i++ {
		if arr[i] > LB(dqueue) {
			for arr[i] > LB(dqueue) {
				dqueue.PopBack()
				if dqueue.LookBack() == nil {
					break
				}
			}
		}
		dqueue.PushBack(arr[i])
	}
	res := []int{LF(dqueue)}
	for i := k; i < n; i++ {
		if arr[i-k] == LF(dqueue) {
			dqueue.PopFront()
		}
		if arr[i] > LB(dqueue) {
			for arr[i] > LB(dqueue) {
				dqueue.PopBack()
				if dqueue.LookBack() == nil {
					break
				}
			}
		}
		dqueue.PushBack(arr[i])
		res = append(res, LF(dqueue))
	}
	return res
}

func LF(deq Struct.DQueue) int {
	return deq.LookFront().(int)
}

func LB(deq Struct.DQueue) int {
	return deq.LookBack().(int)
}

func minSlidingWindow(arr []int, k int) []int {
	if k == 1 {
		return arr
	}
	n := len(arr)
	dqueue := Struct.DQueue{}
	dqueue.PushFront(arr[0])
	for i := 1; i < k; i++ {
		if arr[i] < LB(dqueue) {
			for arr[i] < LB(dqueue) {
				dqueue.PopBack()
				if dqueue.LookBack() == nil {
					break
				}
			}
		}
		dqueue.PushBack(arr[i])
	}
	res := []int{LF(dqueue)}
	for i := k; i < n; i++ {
		if arr[i-k] == LF(dqueue) {
			dqueue.PopFront()
		}
		if arr[i] < LB(dqueue) {
			for arr[i] < LB(dqueue) {
				dqueue.PopBack()
				if dqueue.LookBack() == nil {
					break
				}
			}
		}
		dqueue.PushBack(arr[i])
		res = append(res, LF(dqueue))
	}
	return res
}

func main() {
	arr := []int{5, 5, 2, 1, 3, 10, 0}
	k := 2
	fmt.Println(maxSlidingWindow(arr, k))
	fmt.Println(minSlidingWindow(arr, k))
}
