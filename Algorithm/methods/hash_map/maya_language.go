package main

import "fmt"

func main() {
	w := "sdorg"
	s := "dsrgosdorg"
	fmt.Println(cntWordMaya(w, s))
}

func cntWordMaya(w string, s string) int {
	m, n := len(w), len(s)
	mask := [27]int{}
	wind := [27]int{}
	for i := range w {
		if w[i] == ' ' {
			mask[26]++
		} else {
			mask[w[i]-'a']++
		}
		if s[i] == ' ' {
			wind[26]++
		} else {
			wind[s[i]-'a']++
		}
	}
	res := 0
	eq := 0
	for i := range 26 {
		if wind[i] == mask[i] {
			eq++
		}
	}

	index := -1
	if eq == 26 {
		res++
		index = m - 1
	}

	for i := m; i < n; i++ {
		prev := s[i-m] - 'a'
		if s[i-m] == ' ' {
			prev = 26
		}
		if mask[prev] == wind[prev] {
			eq--
		}
		wind[prev]--
		if mask[prev] == wind[prev] {
			eq++
		}

		post := s[i] - 'a'
		if s[i] == ' ' {
			post = 26
		}
		if mask[post] == wind[post] {
			eq--
		}
		wind[post]++
		if mask[post] == wind[post] {
			eq++
		}

		if eq == 26 && i-m >= index {
			res++
			index = i
		}
	}

	return res
}
