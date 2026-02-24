package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	nums := []int{-10, -9, -8, -5, -4, -1, 0, 1, 2, 4, 6, 8, 9, 10, 11, 14}
	nums2 := []int{-10, -9, -8, -5, -4, -1, 0, 1, 2, 4, 6, 8, 9, 10, 11}
	fmt.Println(compareIntervals(nums))
	fmt.Println(compareIntervals(nums2))
}

func ifBellowZero(n int) string {
	if n < 0 {
		return "(" + strconv.Itoa(n) + ")"
	}
	return strconv.Itoa(n)
}

func compareIntervals(nums []int) string {
	res := strings.Builder{}
	n := len(nums)

	if n < 2 {
		return ""
	}

	l, r := 0, 1
	for ; r < n; r++ {
		if nums[r]-nums[r-1] != 1 {
			res.WriteString(ifBellowZero(nums[l]))
			if r-l > 1 {
				res.WriteByte('-')
				res.WriteString(ifBellowZero(nums[r-1]))
			}
			res.WriteByte(',')
			l = r
		}
	}

	if nums[r-1]-nums[r-2] == 1 {
		res.WriteString(ifBellowZero(nums[l]))
		res.WriteByte('-')
	}

	res.WriteString(ifBellowZero(nums[r-1]))
	res.WriteByte(',')

	return res.String()[0 : len(res.String())-1]
}
