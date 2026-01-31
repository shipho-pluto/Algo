package main

import "fmt"

func MayaLanguage(s, w string) int {
	mask := [27]int{}
	window := [27]int{}
	l1, l2 := len(s), len(w)
	for i := range l2 {
		if w[i] != ' ' {
			mask[w[i]-'a']++
		} else {
			mask[26]++
		}
		if s[i] != ' ' {
			window[s[i]-'a']++
		} else {
			window[26]++
		}
	}

	eq, cnt := 0, 0
	for i := range 27 {
		if mask[i] == window[i] {
			eq++
		}
	}
	flag := false
	if eq == 27 {
		cnt++
		flag = true
	}
	index := 0
	for i := l2; i < l1; i++ {
		prev := s[i-l2] - 'a'
		if s[i-l2] == ' ' {
			prev = 26
		}
		if window[prev] == mask[prev] {
			eq--
		}
		window[prev]--
		if window[prev] == mask[prev] {
			eq++
		}

		post := s[i] - 'a'
		if s[i] == ' ' {
			post = 26
		}
		if window[post] == mask[post] {
			eq--
		}
		window[post]++
		if window[post] == mask[post] {
			eq++
		}
		if eq == 27 {
			if !flag || (i+1+index)%l2 == 0 && (i+index+1) != l2 {
				cnt++
				index = i
				flag = true
			}
		} else {
			flag = false
		}
	}
	return cnt
}

func main() {
	w := "sdorg"
	s := "dsrgosdorg"
	fmt.Println(MayaLanguage(s, w))
}
