package main

import "fmt"

func main() {
	nums := []int{1, 0, 3, 4, 3, 2, 5, 9, 10, 1, 0}
	l, r := []int{0, 4, 6}, []int{2, 7, 8}
	fmt.Println(changesByAdding1(nums, l, r))
}

func changesByAdding1(nums []int, l []int, r []int) []int {
	n := len(nums)
	pref := make([]int, n+1)
	li, ri := 0, 0
	for i := 0; i < n; i++ {
		if li < len(l) && i == l[li] {
			pref[i] += 1
			li++
		}
		if ri < len(r) && i == r[ri] {
			pref[i] -= 1
			ri++
		}
		pref[i+1] += pref[i]
		nums[i] += pref[i]
	}
	return nums
}
