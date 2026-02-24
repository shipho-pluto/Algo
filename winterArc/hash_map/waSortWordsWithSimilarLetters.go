package main

import "fmt"

func main() {
	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(sortWords(arr))
}

func sortWords(arr []string) [][]string {
	st := make(map[[26]int][]string)
	for i := range arr {
		mask := [26]int{}
		for j := range arr[i] {
			mask[arr[i][j]-'a']++
		}
		st[mask] = append(st[mask], arr[i])
	}

	var res [][]string
	for _, v := range st {
		res = append(res, v)
	}

	return res
}
