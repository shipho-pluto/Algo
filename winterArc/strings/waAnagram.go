package main

import "fmt"

func main() {
	s := "anagram"
	t := "nagaran"
	fmt.Println(anagram(s, t))
}

func anagram(s string, t string) bool {
	if len(t) != len(s) {
		return false
	}
	dict := [26]int{}
	for i := range s {
		dict[s[i]-'a']++
		dict[t[i]-'a']--
	}

	for i := range 26 {
		if dict[i] != 0 {
			return false
		}
	}
	return true
}
