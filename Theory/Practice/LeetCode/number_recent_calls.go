package main

import "fmt"

func numberOfRecentCalls(arr [][]int) []int {
	n := len(arr)
	var nums, res []int
	l := 0
	for i := 0; i < n; i++ {
		if arr[i][0] == 1 {
			nums = append(nums, arr[i][2])
		} else {
			cnt := 0
			for t := l; t < len(nums); t++ {
				if nums[t] >= arr[i][1]-3000 {
					cnt++
				} else {
					l = t
				}
			}
			res = append(res, cnt)
		}
	}
	return res
}

func main() {
	arr := [][]int{{1, 1, 1}, {0, 1}, {1, 100, 100}, {0, 100}, {1, 3001, 3001}, {0, 3001}, {1, 3002, 3002}, {0, 3002}}
	fmt.Println(numberOfRecentCalls(arr))
}
