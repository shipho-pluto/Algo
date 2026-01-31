package main

import "fmt"

func main() {
	s, t := "abcabcddredfrty", "fr"
	fmt.Println(findIndexSubstring(s, t))
}

func findIndexSubstring(s, t string) int {
	n, m := len(s), len(t)
	mask, window := [26]int{}, [26]int{}
	for i := range m {
		mask[t[i]-'a']++
		window[s[i]-'a']++
	}
	eq := 0
	for i := range 26 {
		if mask[i] == window[i] {
			eq++
		}
	}
	if eq == 26 {
		return 0
	}
	for i := m; i < n; i++ {
		prev := s[i-m] - 'a'
		if window[prev] == mask[prev] {
			eq--
		}
		window[prev]--
		if window[prev] == mask[prev] {
			eq++
		}

		post := s[i] - 'a'
		if window[post] == mask[post] {
			eq--
		}
		window[post]++
		if window[post] == mask[post] {
			eq++
		}

		if eq == 26 {
			return i - m + 1
		}
	}

	return -1
}
