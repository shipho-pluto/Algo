package main

import "fmt"

func main() {
	arr := []int{1, 2, 2, 1, 0, 2, 2, 2, 0, 2, 1, 2, 1, 0, 1, 1, 2, 2, 1}
	fmt.Println(findNearElem(arr))
}

func findNearElem(arr []int) int {
	n := len(arr)
	l := 0
	index := n
	for ; l < n; l++ {
		if arr[l] == 0 {
			break
		} else if arr[l] == 1 {
			index = min(l, index)
		}
	}
	res := 0
	if index != n {
		res = l - index
	}
	for r := l + 1; r < n; r++ {
		isIn := false
		for ; r < n; r++ {
			if arr[r] == 0 {
				break
			} else if arr[r] == 1 {
				isIn = true
			}
		}

		if isIn {
			for i := l + 1; i < r; i++ {
				if arr[i] == 1 {
					var minimum int
					if r != n {
						minimum = min(i-l, r-i)
					} else {
						minimum = i - l
					}
					res = max(res, minimum)
				}
			}

		}
		l = r
	}
	return res
}
