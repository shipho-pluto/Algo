package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaram"
	fmt.Println(anagramString(s, t))
}

func anagramString(s string, t string) bool {
	m1 := [26]int{}
	n, m := len(s), len(t)
	if n != m {
		return false
	}
	for i := 0; i < n; i++ {
		m1[s[i]-'a']++
		m1[t[i]-'a']--
	}
	for i := range 26 {
		if m1[i] != 0 {
			return false
		}
	}

	return true
}
