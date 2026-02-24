package main

import "fmt"

func main() {
	nums := []int{1, 1, 1, 2, 3, 5, 5, 5}
	k := 4
	fmt.Println(mostRepeatBucket(nums, k))
}

func mostRepeatBucket(arr []int, k int) []int {
	cnt := make(map[int]int)
	for i := range arr {
		cnt[arr[i]]++
	}

	b := make([][]int, len(arr))
	for i, v := range cnt {
		b[v-1] = append(b[v-1], i)
	}

	res := make([]int, 0, k)
	for i := len(b) - 1; i > -1; i-- {
		for j := range len(b[i]) {
			res = append(res, b[i][j])
			k--
			if k == 0 {
				return res
			}
		}
	}

	return []int{}
}
