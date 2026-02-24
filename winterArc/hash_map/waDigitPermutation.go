package main

import "fmt"

func main() {
	arr := []int{1230, 99, 23001, 123, 111, 300021, 101010, 90000009, 9}
	fmt.Println(digitPerm(arr))
}

func digitPerm(arr []int) [][]int {
	st := make(map[[9]int][]int)

	for i := range arr {
		mask := [9]int{}
		c := arr[i]
		for arr[i] != 0 {
			if arr[i]%10 != 0 {
				mask[arr[i]%10-1]++
			}
			arr[i] /= 10
		}
		st[mask] = append(st[mask], c)
	}

	var res [][]int
	for _, v := range st {
		res = append(res, v)
	}

	return res
}
