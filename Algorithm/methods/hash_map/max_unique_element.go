package main

import "fmt"

func findMaxUniqueElement(arr string) int {
	idx := make(map[string]int)
	p := 0
	length := 0
	for i := 0; i < len(arr); i++ {
		p = max(p, idx[string(arr[i])])
		length = max(length, i-p+1)
		idx[string(arr[i])] = i + 1
	}

	return length
}

func main() {
	s := "abcabcabcdabcqwerty"
	fmt.Println(findMaxUniqueElement(s))
}
