package main

import "fmt"

func twoSum(nums []int, target int) []int {
	pref := make(map[int]int)
	pref[0] = 0
	for i := 0; i < len(nums); i++ {
		if val, ok := pref[target-nums[i]]; ok {
			return []int{val, i}
		}
		pref[nums[i]] = i
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 17
	fmt.Println(twoSum(nums, target))
}
