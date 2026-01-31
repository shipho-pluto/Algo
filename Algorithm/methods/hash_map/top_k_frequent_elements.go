package main

import "fmt"

func topKFrequent(nums []int, k int) []int {
	cnt := make(map[int]int)
	for _, num := range nums {
		cnt[num]++
	}

	bucket := make([][]int, len(nums)+1)
	for key, val := range cnt {
		bucket[val] = append(bucket[val], key)
	}

	res := make([]int, k)
LOOP:
	for i := len(bucket) - 1; i > -1; i-- {
		for j := range bucket[i] {
			if k > 0 {
				res[k-1] = bucket[i][j]
				k--
			} else {
				break LOOP
			}
		}
	}

	return res
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3, 5, 5, 5}
	k := 2
	fmt.Println(topKFrequent(nums, k))
}
