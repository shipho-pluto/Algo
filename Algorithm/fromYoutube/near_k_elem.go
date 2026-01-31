package main

import "fmt"

func main() {
	arr := []int{-1, 0, 2, 3, 5, 6, 9}
	index, k := 2, 5
	fmt.Println(nearKElem(arr, index, k))
}

func nearKElem(arr []int, index int, k int) []int {
	n := len(arr)
	if index >= n || k > n || k == 0 {
		return []int{}
	}

	l, r := index, index+1
	res := make([]int, k)
	cnt := k
	cur := arr[index]
	for i := 0; (l >= 0 || r < n) && cnt > 0; cnt, i = cnt-1, i+1 {
		if l >= 0 && r < n {
			if cur-arr[l] < arr[r]-cur {
				res[i] = arr[l]
				l--
			} else {
				res[i] = arr[r]
				r++
			}
		} else if l >= 0 {
			res[i] = arr[l]
			l--
		} else {
			res[i] = arr[r]
			r++
		}
	}

	return res
}
