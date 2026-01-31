package main

import "fmt"

func Abs1(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func FindKNearToI(arr []int, k, index int) []int {
	r, l := index+1, index
	n := len(arr)
	var res []int
	for l > -1 || r < n {
		if l > -1 && r < n {
			if Abs1(arr[l]-arr[index]) < Abs1(arr[r]-arr[index]) {
				res = append(res, arr[l])
				l--
			} else {
				res = append(res, arr[r])
				r++
			}
		} else if r < n {
			res = append(res, arr[r])
			r++
		} else if l > -1 {
			res = append(res, arr[l])
			l--
		}
		if len(res) == k {
			break
		}
	}
	return res
}

func main() {
	arr := []int{-2, -1, 0, 4, 10, 15, 19}
	fmt.Println(FindKNearToI(arr, 4, 4))
}
