package main

import "fmt"

func main() {
	s := "hello, world!"
	for _, v := range printHistogram(s) {
		fmt.Println(v)
	}
}

func printHistogram(s string) []string {
	n := len(s)
	mask := make(map[uint8]int, n)
	maxC := 0
	for i := range n {
		mask[s[i]]++
		maxC = max(maxC, mask[s[i]])
	}

	st := make([][]int, len(mask))
	idx := 0
	for i := range mask {
		st[idx] = []int{int(i), mask[i]}
		idx++
	}

	res := make([]string, maxC+1)
	for i := range st {
		res[0] += string(rune(st[i][0]))
	}
	for i := range maxC {
		for k := range st {
			if st[k][1] > 0 {
				st[k][1]--
				res[i+1] += "#"
			} else {
				res[i+1] += " "
			}
		}
	}

	result := make([]string, maxC+1)
	for l, r := 0, maxC; r > -1; l, r = l+1, r-1 {
		result[l] = res[r]
	}
	return result
}
