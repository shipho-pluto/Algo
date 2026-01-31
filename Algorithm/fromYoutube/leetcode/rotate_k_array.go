package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(arr, k)
	fmt.Println(arr)
}

func rotate(nums []int, k int) {
	k %= len(nums)
	pref := make([]int, k)
	copy(pref, nums[len(nums)-k:])
	for i := len(nums) - 1; i > k-1; i-- {
		nums[i] = nums[i-k]
	}
	copy(nums, pref)
}
