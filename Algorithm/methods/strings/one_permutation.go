package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}

func checkInclusion(s1 string, s2 string) bool {
	mask := [26]int{}
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	for i := range n {
		mask[s1[i]-'a']++
		mask[s2[i]-'a']--
	}
	eq := 0
	for i := range 26 {
		if mask[i] == 0 {
			eq++
		}
	}
	if eq == 26 {
		return true
	}
	for i := n; i < m; i++ {
		prev := s2[i-n] - 'a'
		if mask[prev] == 0 {
			eq--
		}
		mask[prev]++
		if mask[prev] == 0 {
			eq++
		}

		post := s2[i] - 'a'
		if mask[post] == 0 {
			eq--
		}
		mask[post]--
		if mask[post] == 0 {
			eq++
		}

		if eq == 26 {
			return true
		}
	}
	return false
}
