package main

import "fmt"

func cntUniqueElements(s string) int {
	n := len(s)
	mask := make(map[string]bool)
	idx := 0
	for ; idx < n; idx++ {
		if mask[string(s[idx])] {
			break
		}
		mask[string(s[idx])] = true
	}

	res := idx * (idx - 1) / 2
	l := 0
	for r := idx; r < n; r++ {
		if mask[string(s[r])] {
			for ; l < n; l++ {
				if !mask[string(s[r])] {
					break
				}
				mask[string(s[l])] = false
			}
		}
		res += r - l
		mask[string(s[r])] = true
	}

	return res
}

func main() {
	s := "aabbcdb"
	fmt.Println(cntUniqueElements(s))
}
