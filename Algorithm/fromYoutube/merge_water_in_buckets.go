package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 3, 2, 4, 10, 8, 4, 2, 5, 3}
	buckets := 4
	fmt.Println(mergeWaterInBuckets(arr, buckets))
}

func check(arr []int, bucketSum int, cntBuck int) bool {
	sum := 0
	cnt := 0
	for i := range arr {
		if arr[i]+sum > bucketSum {
			sum = 0
			cnt++
		}
		sum += arr[i]
	}
	if sum != 0 {
		cnt++
	}
	return cnt <= cntBuck
}

func minArr(arr []int) int {
	m := arr[0]
	for i := range arr {
		if arr[i] > m {
			m = arr[i]
		}
	}
	return m
}

func maxArr(arr []int) int {
	m := 0
	for i := range arr {
		m += arr[i]
	}
	return m
}

func mergeWaterInBuckets(arr []int, cntBuckets int) int {
	l, r := minArr(arr), maxArr(arr)
	for l < r {
		m := (l + r) / 2
		if check(arr, m, cntBuckets) {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
