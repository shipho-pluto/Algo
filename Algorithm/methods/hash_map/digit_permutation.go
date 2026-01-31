package main

import "fmt"

func main() {
	arr := []int{1230, 99, 23001, 123, 111, 300021, 101010, 90000009, 9}
	fmt.Println(digitPermutation(arr))
}

func digitPermutation(arr []int) [][]int {
	n := len(arr)
	st := make(map[[9]int][]int, n)
	for i := 0; i < n; i++ {
		mask := [9]int{}
		el := arr[i]
		for arr[i] > 0 {
			if arr[i]%10 != 0 {
				mask[arr[i]%10-1]++
			}
			arr[i] /= 10
		}
		st[mask] = append(st[mask], el)
	}

	var res [][]int
	for _, val := range st {
		res = append(res, val)
	}
	return res
}
