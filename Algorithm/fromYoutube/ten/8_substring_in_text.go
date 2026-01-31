package main

import "fmt"

func main() {
	s, t := "abcabcdfr", "fcd"
	fmt.Println(substringInText(t, s))
}

func substringInText(t string, s string) int {
	n, m := len(s), len(t)
	mask, wind := [26]int{}, [26]int{}
	for i := range m {
		mask[t[i]-'a']++
		wind[s[i]-'a']++
	}
	eq := 0
	for i := range 26 {
		if mask[i] == wind[i] {
			eq++
		}
	}
	if eq == 26 {
		return 0
	}
	for i := m; i < n; i++ {
		prev := s[i-m] - 'a'
		if mask[prev] == wind[prev] {
			eq--
		}
		wind[prev]--
		if mask[prev] == wind[prev] {
			eq++
		}

		post := s[i] - 'a'
		if mask[post] == wind[post] {
			eq--
		}
		wind[post]++
		if mask[post] == wind[post] {
			eq++
		}

		if eq == 26 {
			return i - m + 1
		}
	}
	return -1
}
