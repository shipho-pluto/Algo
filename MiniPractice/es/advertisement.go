package main

import (
	"fmt"
	"sort"
)

func CalculateAdvertisement(arr [][]int, k int) (int, int) {
	n := len(arr)
	var st [][]int
	for i := 0; i < n; i++ {
		if arr[i][1]-arr[i][0] >= k {
			st = append(st, []int{arr[i][0], 1})
			st = append(st, []int{arr[i][1], -1})
		}
	}

	sort.Slice(st, func(i, j int) bool {
		if st[i][0] == st[j][0] {
			return st[i][1] < st[j][1]
		}
		return st[i][0] < st[j][0]
	})

	prefix := [][]int{{0, 0}}
	cnt := 0
	for _, v := range st {
		cnt += v[1]
		prefix = append(prefix, []int{v[0], cnt})
	}

	res := [][]int{{0, 0}}
	index := 0
	for i := range prefix {
		if prefix[i][1] == 0 {
			res = append(res, []int{0, 0})
			index++
		}
		if res[index][1] < prefix[i][1] {
			res[index] = []int{prefix[i][0], prefix[i][1]}
		}
	}

	max1, max2 := []int{0, 0}, []int{0, 0}
	for i := range res {
		if res[i][1] > max1[1] {
			max2 = max1
			max1 = res[i]
		} else {
			if res[i][1] > max2[1] {
				max2 = res[i]
			}
		}
	}
	return max2[0], max1[0]
}

func main() {
	n := [][]int{{1, 4}, {4, 6}, {5, 6}, {8, 10}, {6, 9}, {7, 9}, {11, 17}, {10, 14}, {11, 15}, {12, 13}}
	fmt.Println(CalculateAdvertisement(n, 1))
}
