package main

import "fmt"

func findPreString(t, s string) int {
	mask := [26]int{}
	find := [26]int{}
	n, m := len(t), len(s)
	for i := 0; i < m; i++ {
		mask[s[i]-'a']++
		find[t[i]-'a']++
	}

	eq := 0
	for i := range 26 {
		if mask[i] == find[i] {
			eq++
		}
	}
	if eq == 26 {
		return 0
	}
	for i := m; i < n; i++ {
		pred := t[i-m] - 'a'
		if mask[pred] == find[pred] {
			eq--
		}
		find[pred]--
		if mask[pred] == find[pred] {
			eq++
		}

		post := t[i] - 'a'
		if mask[post] == find[post] {
			eq--
		}
		find[post]++
		if mask[post] == find[post] {
			eq++
		}

		if eq == 26 {
			return i - m + 1
		}
	}
	return -1
}

func main() {
	t := "abcabcaabbcc"
	s := "ccb"
	fmt.Println(findPreString(t, s))
}
