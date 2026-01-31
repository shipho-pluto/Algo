package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{-3, -2, -1, 0, 0, 1, 2, 3}
	target := 0
	fmt.Println(fourSum(arr, target))
}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	n := len(nums)
	var res [][]int
	var flag, signal bool
	for l, r := 0, n-1; l < r; {
		if flag {
			flag = false
		} else {
			flag = true
		}
		li, ri := l+1, r-1
		for li < ri {
			if nums[li]+nums[ri]+nums[l]+nums[r] == target {
				res = append(res, []int{nums[l], nums[li], nums[ri], nums[r]})
				li++
				for li < ri && nums[li] == nums[li-1] {
					li++
				}
				ri--
				for li < ri && nums[ri] == nums[ri+1] {
					ri--
				}
			} else if nums[li]+nums[ri]+nums[l]+nums[r] < target {
				li++
			} else {
				ri--
			}
		}

		if signal || nums[l]+nums[l+1]+nums[r-1]+nums[r] == target {
			if flag {
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
				if r < n-1 {
					r++
				}
			} else {
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
				if l > 0 {
					l--
				}
			}
			signal = true
		} else if nums[l]+nums[l+1]+nums[r-1]+nums[r] < target {
			l++
			for l < r && nums[l] == nums[l-1] {
				l++
			}
			signal = false
		} else {
			r--
			for l < r && nums[r] == nums[r+1] {
				r--
			}
			signal = false
		}
	}

	return res
}
