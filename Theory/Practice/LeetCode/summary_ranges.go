package main

import "fmt"

func summaryRanges(arr []int) []int {
	n := len(arr)
	cnt := 0
	var res []int
	j := 0
	for ; j < n-1; j++ {
		cnt += arr[j]
		if arr[j]+1 != arr[j+1] {
			res = append(res, cnt)
			cnt = 0
		}
	}

	if arr[j-1]+1 != arr[j] {
		res = append(res, arr[j])
	} else {
		res = append(res, cnt+arr[j])
	}

	return res
}

func main() {
	arr := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(arr))
}
