package main

import (
	"fmt"
	"sort"
)

func Hotel(arr [][]int) int {
	n := len(arr)
	res := make([][]int, 2*n)
	index := 0
	for i := range arr {
		res[index] = []int{arr[i][0], 1}
		res[index+1] = []int{arr[i][1], -1}
		index += 2
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i][0] == res[j][0] {
			return res[i][1] < res[j][1]
		}
		return res[i][0] < res[j][0]
	})

	cnt := 0
	maxCnt := 0
	for i := range res {
		cnt += res[i][1]
		maxCnt = max(maxCnt, cnt)
	}

	return maxCnt
}

func main() {
	arr := [][]int{{0, 2}, {2, 3}, {3, 4}, {0, 4}}

	fmt.Println(Hotel(arr))
}
