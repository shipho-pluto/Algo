package main

import "fmt"

func main() {
	arr1, arr2 := []int{4, 9, 5}, []int{9, 4, 9, 8, 4}
	fmt.Println(interceptionTwoArray(arr1, arr2))

	s1, s2 := "abbcde", "ab"
	fmt.Println(interceptionTwoString(s1, s2))
}

func interceptionTwoString(s1 string, s2 string) bool {
	hm := [26]int{}
	for i := range s1 {
		hm[s1[i]-'a']++
	}
	for i := range s2 {
		if val := hm[s2[i]-'a']; val > 0 {
			hm[s2[i]-'a']--
		} else {
			return false
		}
	}
	return true
}

func interceptionTwoArray(arr1 []int, arr2 []int) []int {
	var res []int
	hm := make(map[int]int)
	for i := range arr1 {
		hm[arr1[i]]++
	}
	for i := range arr2 {
		if _, ok := hm[arr2[i]]; ok {
			res = append(res, arr2[i])
			hm[arr2[i]]--
			if hm[arr2[i]] == 0 {
				delete(hm, arr2[i])
			}
		}
	}
	return res
}
