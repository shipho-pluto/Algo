package main

import "fmt"

func main() {
	arr := []int{1, 2, 0, 1, 1, 1, 0, 0, 0, 0, 0, 2, 0, 1, 1, 1, 0, 1}
	fmt.Println(maxDistanceCloset(arr))
}

func maxDistanceCloset(arr []int) int {
	n := len(arr)
	l := 0
	for ; l < n; l++ {
		if arr[l] == 1 {
			break
		}
	}
	res := l

	for r := l + 1; r < n; r++ {
		exist := false
		for ; r < n; r++ {
			if arr[r] == 0 {
				exist = true
			} else if arr[r] == 1 {
				break
			}
		}

		if exist {
			for i := l + 1; i < r; i++ {
				if arr[i] == 0 {
					if r == n {
						res = max(res, i-l)
					} else {
						res = max(res, min(i-l, r-i))
					}
				}
			}
		}
		l = r
	}

	return res
}
