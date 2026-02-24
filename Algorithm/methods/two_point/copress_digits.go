package main

import (
	"fmt"
	"sort"
	"strconv"
)

func compressDigits(arr []int) string {
	n := len(arr)
	if n == 0 {
		return ""
	}
	sort.Ints(arr)

	res := ""
	var prev int

	for l, r := 0, 1; r < n; l, r = l+1, r+1 {
		prev = l
		for ; r < n && arr[r] == arr[l]+1; l, r = l+1, r+1 {
		}
		res += strconv.Itoa(arr[prev])
		if r-prev == 1 {
			res += ","
		} else {
			res += "-" + strconv.Itoa(arr[r-1]) + ","
		}
	}

	if arr[n-1] != arr[n-2]+1 {
		res += strconv.Itoa(arr[n-1]) + ","
	}

	return res[:len(res)-1]
}

func main() {
	arr := []int{-10, -9, -5, -1, 0, 1, 3, 4, 6, 7, 8, 11, 12, 13, 15}
	fmt.Println(compressDigits(arr))
}
