package main

import "fmt"

func interceptionTwoArrays(arr1, arr2 []int) []int {
	n, m := len(arr1), len(arr2)
	st := make(map[int]int)
	for i := range n {
		st[arr1[i]]++
	}
	var res []int
	for i := range m {
		if _, ok := st[arr2[i]]; ok {
			res = append(res, arr2[i])
			st[arr2[i]]--
			if st[arr2[i]] == 0 {
				delete(st, arr2[i])
			}
		}
	}
	return res
}

func interceptionTwoStrings(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	st := make(map[string]int)
	for i := range n {
		st[string(s1[i])]++
	}
	for i := range m {
		if _, ok := st[string(s2[i])]; ok {
			st[string(s2[i])]--
			if st[string(s2[i])] < 0 {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func main() {
	arr1 := []int{4, 9, 5}
	arr2 := []int{9, 4, 9, 8, 4}
	fmt.Println(interceptionTwoArrays(arr1, arr2))

	s1 := "ahngftt"
	s2 := "ahtxt"
	fmt.Println(interceptionTwoStrings(s1, s2))
}
