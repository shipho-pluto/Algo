package main

import "fmt"

func main() {
	s1, s2 := "cat", "tcat"
	fmt.Println(oneEdit(s1, s2))
}

func oneEdit(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		m, n = n, m
		s1, s2 = s2, s1
	}
	if m-n > 1 {
		return false
	}

	if n == m {
		neq := 0
		for i := range n {
			if s1[i] != s2[i] {
				neq++
			}
			if neq == 2 {
				return false
			}
		}
	} else {
		for i := range n {
			if s1[i] != s2[i] && s1[i] != s2[i+1] {
				return false
			}
		}
	}
	return true
}
