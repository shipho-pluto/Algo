package main

import "fmt"

func main() {
	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(sortWords(arr))
}

func sortWords(arr []string) [][]string {
	var res [][]string
	n := len(arr)
	st := make(map[[26]int][]string)
	for i := range n {
		mask := [26]int{}
		for j := range arr[i] {
			mask[arr[i][j]-'a']++
		}
		st[mask] = append(st[mask], arr[i])
	}
	for _, val := range st {
		res = append(res, val)
	}
	return res
}
