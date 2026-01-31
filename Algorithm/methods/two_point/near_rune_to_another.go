package main

import "fmt"

func nearXToY(s string) int {
	var xXORy bool
	index := 0
	n := len(s)

	for i := 0; i < n; i++ {
		if s[i] == 'X' {
			xXORy = true
			index = i
			break
		} else if s[i] == 'Y' {
			xXORy = false
			index = i
			break
		}
	}

	minLen := n
	for i := index + 1; i < n; i++ {
		if s[i] == 'Y' && xXORy {
			minLen = min(i-index, minLen)
			index = i
			xXORy = false
		} else if s[i] == 'Y' && !xXORy {
			index = i
		} else if s[i] == 'X' && !xXORy {
			minLen = min(i-index, minLen)
			index = i
			xXORy = true
		} else if s[i] == 'X' && xXORy {
			index = i
		}
	}

	if minLen < n {
		return minLen
	}
	return 0
}

func main() {
	ar := "000XOOYOX0"
	fmt.Println(nearXToY(ar))
}
