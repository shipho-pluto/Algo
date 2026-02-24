package main

import "fmt"

func main() {
	temple := "Hello, World!"
	res := histogram(temple)
	for i := range res[0] {
		for j := range res {
			fmt.Print(string(res[j][len(res[0])-1-i]))
		}
		fmt.Println()
	}
}

func histogram(t string) []string {
	cnt := make(map[uint8]int)
	maxL := 0
	for i := range t {
		cnt[t[i]]++
		maxL = max(maxL, cnt[t[i]])
	}
	res := make([]string, len(cnt))
	i := 0
	for k, v := range cnt {
		res[i] += string(k)
		for j := range maxL {
			if j < v {
				res[i] += "#"
			} else {
				res[i] += " "
			}
		}
		i++
	}
	return res
}
