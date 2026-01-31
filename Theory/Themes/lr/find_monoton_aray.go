package main

import "fmt"

func maxMonotonic(arr []int) (int, int) {
	n := len(arr)
	if n == 1 {
		return 0, 0
	}
	l, r := 0, 1
	cnt := 1
	monotonic := true
	res := []int{1, 0, 0}
	for r < n {
		if arr[r-1] < arr[r] && monotonic {
			r++
			cnt++
		} else if arr[r-1] > arr[r] && monotonic {
			if cnt > res[0] {
				res = []int{cnt, l, r - 1}
			}
			l = r - 1
			cnt = 1
			monotonic = false
		} else if arr[r-1] > arr[r] && !monotonic {
			r++
			cnt++
		} else if arr[r-1] < arr[r] && !monotonic {
			if cnt > res[0] {
				res = []int{cnt, l, r - 1}
			}
			l = r - 1
			cnt = 1
			monotonic = true
		} else if arr[r-1] == arr[r] {
			if cnt > res[0] {
				res = []int{cnt, l, r - 1}
			}
			l = r
			r++
			cnt = 1
		}
	}

	if arr[r-2] < arr[r-1] && monotonic || arr[r-2] > arr[r-1] && !monotonic {
		if cnt > res[0] {
			res = []int{cnt, l, r - 1}
		}
	}

	return res[1], res[2]
}

func main() {
	arr := []int{2}
	fmt.Println(maxMonotonic(arr))
}
