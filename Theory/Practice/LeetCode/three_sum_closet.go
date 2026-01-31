package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{2, 3, 8, 9, 10}
	target := 16
	res := threeSumClosest(arr, target)

	fmt.Println(res)
}

func threeSumClosest(arr []int, target int) int {
	n := len(arr)
	sort.Ints(arr)
	minSum := arr[0] + arr[1] + arr[2]
	minDif := target - minSum
	for i := 0; i < n-2; i++ {
		if i >= 1 && arr[i] != arr[i-1] || i == 0 {
			for l, r := i+1, n-1; l < r; {
				if arr[i]+arr[l]+arr[r] == target {
					return target
				} else if arr[i]+arr[l]+arr[r] < target {
					if minDif > target-(arr[i]+arr[l]+arr[r]) {
						minDif = target - arr[i] + arr[l] + arr[r]
						minSum = arr[i] + arr[l] + arr[r]
					}
					l++
				} else {
					if minDif > arr[i]+arr[l]+arr[r]-target {
						minDif = arr[i] + arr[l] + arr[r] - target
						minSum = arr[i] + arr[l] + arr[r]
					}
					r--
				}
			}
		}
	}
	return minSum
}
