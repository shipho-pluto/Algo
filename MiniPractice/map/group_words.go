package main

import "fmt"

func GroupWords(arr []string) [][]string {
	maskAll := map[[26]int][]string{}
	n := len(arr)
	for i := 0; i < n; i++ {
		m := len(arr[i])
		mask := [26]int{}
		for j := 0; j < m; j++ {
			mask[arr[i][j]-'a']++
		}
		maskAll[mask] = append(maskAll[mask], arr[i])
	}
	m := len(maskAll)
	res := make([][]string, m)
	index := 0
	for _, val := range maskAll {
		res[index] = val
		index++
	}
	return res
}

func main() {
	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(GroupWords(arr))
}
