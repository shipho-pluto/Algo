package main

import "fmt"

func GO(i int, m map[int]int, nums []int, sum int) int {
	if i+1 >= len(nums) {
		return 0
	}
	m[sum+nums[i+1]] = (m[sum] << 1) + 1
	m[sum-nums[i+1]] = m[sum] << 1
	GO(i+1, m, nums, sum+nums[i+1])
	GO(i+1, m, nums, sum-nums[i+1])
	return 0
}

func ansByNub(n int) string {
	ans := ""
	for n != 0 {
		if n%2 == 1 {
			ans = "+" + ans
		} else {
			ans = "-" + ans
		}
		n /= 2
	}
	return ans[1:]
}

func main() {
	nums2 := []int{0, 0}
	target2 := 0
	m2 := make(map[int]int)
	m2[nums2[0]] = 1
	_ = GO(0, m2, nums2, nums2[0])
	fmt.Println(ansByNub(m2[target2]))
	fmt.Println(m2)
}
